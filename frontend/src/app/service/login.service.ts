import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpStatusCode } from '@angular/common/http';
import { catchError, lastValueFrom, map, Observable, of, Subject, tap } from 'rxjs';
import { base64URLStringToBuffer, bufferToBase64URLString } from '../util/util';
import {
  AUTH_STATE,
  AuthenticationCredential,
  IFinalizeLoginResponse,
  ILoginData, IToken,
  WrappedPublicKeyCredentialRequestOptionsJSON,
} from '../model/types';
import { SnackbarService } from './snackbar.service';
import { AppStateService } from './appstate.service';
import jwtDecode from 'jwt-decode';


@Injectable({
  providedIn: 'root',
})
export class LoginService {
  constructor(
    private httpClient: HttpClient,
    private snackbarService: SnackbarService,
    private appStateService: AppStateService,
  ) {}

  public login(email: string): Observable<AUTH_STATE> {
    const resultSub = new Subject<AUTH_STATE>()
    this.httpClient.post<WrappedPublicKeyCredentialRequestOptionsJSON>('https://fido.workshop/api/v1/service/authenticate/begin', {
      email,
    }).pipe(
      tap(res => console.log('Result of https://fido.workshop/api/v1/service/authenticate/begin call', res)),
    ).subscribe({
      next: res => {
        this.handleLoginAttemptResponse(res).then(isValidated => {
          resultSub.next(isValidated);
          resultSub.complete();
        });
      },
      error: (err: HttpErrorResponse) => {
        console.error(`Error while attempting login with email '${email}'`, err);
        resultSub.next(HttpStatusCode.Unauthorized ? AUTH_STATE.UNKNOWN_USER : AUTH_STATE.HTTP_ERROR);
        resultSub.complete();
      }
    });
    return resultSub.asObservable();
  }

  private finalizeLogin(loginData: ILoginData): Observable<boolean> {
    return this.httpClient.post<IFinalizeLoginResponse>('https://fido.workshop/api/v1/service/authenticate/finish', {
      ...loginData,
    }).pipe(
      tap(res => console.log('Result of https://fido.workshop/api/v1/service/authenticate/finish call', res)),
      map(res => {
        if (!res || !res.token) {
          // not logged in
          return false;
        }

        try {
          const token = jwtDecode<IToken>(res.token);

          console.log('token', token);
          this.appStateService.setUser({
            firstName: token.firstName,
            lastName: token.lastName,
            mail: token.mail,
          })
          this.appStateService.setToken(res.token, token);
          this.appStateService.setLoginState(true);
        } catch (err) {
          this.snackbarService.showSnackbarWithMessage('Error while decoding the jwt');
          return false;
        }

        // logged in
        return true;
      }),
      catchError(() => of(false)),
    );
  }

  private async handleLoginAttemptResponse(requestOptionsJSON: WrappedPublicKeyCredentialRequestOptionsJSON): Promise<AUTH_STATE> {
    const options: CredentialRequestOptions = {
      publicKey: {
        ...requestOptionsJSON,
        challenge: base64URLStringToBuffer(requestOptionsJSON.publicKey.challenge),
        allowCredentials: (requestOptionsJSON.publicKey.allowCredentials || []).map(allowCredential => ({
          ...allowCredential,
          id: base64URLStringToBuffer(allowCredential.id),
        })),
      }
    };

    console.log('handleLoginAttemptResponse options:', options);

    const credentials = await (navigator.credentials.get(options) as any as Promise<AuthenticationCredential | null>);

    if (!credentials) {
      console.error('Cannot get credentials via navigator.credentials.get call');
      return AUTH_STATE.AUTHENTICATOR_ERROR;
    }

    console.log('credentials', credentials);

    const isLoginValidated = await lastValueFrom(this.finalizeLogin({
      id: credentials.id,
      rawId: bufferToBase64URLString(credentials.rawId),
      response: {
        authenticatorData: bufferToBase64URLString(credentials.response.authenticatorData),
        clientDataJSON: bufferToBase64URLString(credentials.response.clientDataJSON),
        userHandle: credentials.response.userHandle ? bufferToBase64URLString(credentials.response.userHandle) : '',
        signature: bufferToBase64URLString(credentials.response.signature),
      },
      type: credentials.type,
    }));

    return isLoginValidated ? AUTH_STATE.VERIFIED : AUTH_STATE.NOT_VERIFIED;
  }
}
