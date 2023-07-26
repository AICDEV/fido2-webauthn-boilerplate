import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpStatusCode } from '@angular/common/http';
import { catchError, Observable, of, Subject, tap, throwError } from 'rxjs';
import { base64URLStringToBuffer, bufferToBase64URLString } from '../util/util';
import { AUTH_STATE, IRegisterResponse, RegistrationCredential, RegistrationResponseJSON } from '../model/types';
import { AppStateService } from './appstate.service';


@Injectable({
  providedIn: 'root',
})
export class RegisterService {

  constructor(
    private httpClient: HttpClient,
    private appStateService: AppStateService,
  ) {}

  public register(mail: string, firstName: string, lastName: string): Observable<AUTH_STATE> {
    const result$ = new Subject<AUTH_STATE>();

    if (!window.PublicKeyCredential) {
      alert("Error: this browser does not support WebAuthn");
      return of(AUTH_STATE.AUTHENTICATOR_ERROR);
    }

    // get the server challenge
    this.httpClient.post<IRegisterResponse>('https://fido.workshop/api/v1/service/signup/begin', {
      displayName: `${firstName} ${lastName}`,
      firstName: firstName,
      lastName: lastName,
      email: mail,
    }).pipe(
      tap(res => console.log('Response of https://fido.workshop/api/v1/service/signup/begin call', res)),
    ).subscribe({
      next: response => {
        // create the credentials out of the server data
        this.credentialCreateWithRegisterResponse(response).then(credentialInfo => {
          if (!credentialInfo) {
            console.error('Error on navigator.credentials.create. Cannot get credentialInfo');
            result$.next(AUTH_STATE.AUTHENTICATOR_ERROR);
            result$.complete();
            return;
          }
          console.log('credentialInfo', credentialInfo);

          const registerResponse = this.createFinalRegistrationValidationData(credentialInfo);

          // validate and finalize the signup
          this.httpClient.post('https://fido.workshop/api/v1/service/signup/finish', registerResponse)
            .pipe(
              catchError((err: HttpErrorResponse) => {
                console.error('Signup finish failed', err);
                result$.next(err.status === HttpStatusCode.Unauthorized ? AUTH_STATE.NOT_VERIFIED : AUTH_STATE.HTTP_ERROR);
                result$.complete();
                return throwError(err);
              })
            ).subscribe(() => {
            result$.next(AUTH_STATE.VERIFIED);
            result$.complete();
          });
        }).catch(err => {
          console.log(err);
          result$.next(AUTH_STATE.AUTHENTICATOR_ERROR);
          result$.complete();
        });
      },
      error: (err: HttpErrorResponse) => {
        console.error(`Error while attempting new registration with email '${mail}'`, err);
        result$.next(err.status === HttpStatusCode.BadRequest ? AUTH_STATE.NOT_VERIFIED : AUTH_STATE.HTTP_ERROR);
        result$.complete();
      }
    });

    return result$.asObservable();
  }

  private createFinalRegistrationValidationData(credentialInfo: RegistrationCredential): RegistrationResponseJSON  {
    let responsePublicKey: string | undefined = undefined;
    if (typeof credentialInfo.response.getPublicKey === 'function') {
      const _publicKey = credentialInfo.response.getPublicKey();
      if (_publicKey !== null) {
        responsePublicKey = bufferToBase64URLString(_publicKey);
      }
    }

    // FIREFOX COMPATIBILITY CODE
    let transports: AuthenticatorTransport[] | undefined = undefined;
    if (typeof credentialInfo.response.getTransports === 'function') {
      transports = credentialInfo.response.getTransports() as AuthenticatorTransport[];
    }

    let responsePublicKeyAlgorithm: number | undefined = undefined;
    if (typeof credentialInfo.response.getPublicKeyAlgorithm === 'function') {
      responsePublicKeyAlgorithm = credentialInfo.response.getPublicKeyAlgorithm();
    }

    let authenticatorData: string | undefined = undefined;
    if (typeof credentialInfo.response.getAuthenticatorData === 'function') {
      authenticatorData = bufferToBase64URLString(credentialInfo.response.getAuthenticatorData());
    }
    // FIREFOX COMPATIBILITY CODE END

    return {
        id: credentialInfo.id,
        rawId: bufferToBase64URLString(credentialInfo.rawId),
        response: {
          attestationObject: bufferToBase64URLString(credentialInfo.response.attestationObject),
            clientDataJSON: bufferToBase64URLString(credentialInfo.response.clientDataJSON),
            transports,
            publicKeyAlgorithm: responsePublicKeyAlgorithm,
            publicKey: responsePublicKey,
            authenticatorData: authenticatorData,
        },
        type: 'public-key',
        clientExtensionResults: credentialInfo.getClientExtensionResults(),
        authenticatorAttachment: credentialInfo.authenticatorAttachment ? credentialInfo.authenticatorAttachment as AuthenticatorAttachment: undefined,
    };
  }

  private credentialCreateWithRegisterResponse(registerResponse: IRegisterResponse): Promise<RegistrationCredential | null> {
    return (navigator.credentials.create({
      publicKey: {
        challenge: base64URLStringToBuffer(registerResponse.publicKey.challenge),
        rp: registerResponse.publicKey.rp,
        user: {
          id: base64URLStringToBuffer(registerResponse.publicKey.user.id),
          displayName: registerResponse.publicKey.user.displayName,
          name: registerResponse.publicKey.user.name,
        },
        pubKeyCredParams: registerResponse.publicKey.pubKeyCredParams,
        authenticatorSelection: registerResponse.publicKey.authenticatorSelection,
      },
    })) as Promise<RegistrationCredential | null>;
  }

  public registerNewKey(): Observable<AUTH_STATE> {
    const result$ = new Subject<AUTH_STATE>();

    if (!window.PublicKeyCredential) {
      alert("Error: this browser does not support WebAuthn");
      return of(AUTH_STATE.AUTHENTICATOR_ERROR);
    }

    // get the server challenge
    this.httpClient.post<IRegisterResponse>('https://fido.workshop/api/v1/member/register/device/begin', undefined, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.appStateService.getRawToken()}`,
      },
    }).pipe(
      tap(res => console.log('Response of https://fido.workshop/api/v1/member/register/device/begin call', res)),
    ).subscribe({
      next: response => {
        // create the credentials out of the server data
        this.credentialCreateWithRegisterResponse(response).then(credentialInfo => {
          if (!credentialInfo) {
            console.error('Error on navigator.credentials.create. Cannot get credentialInfo');
            result$.next(AUTH_STATE.AUTHENTICATOR_ERROR);
            result$.complete();
            return;
          }
          console.log('credentialInfo', credentialInfo);

          const registerResponse = this.createFinalRegistrationValidationData(credentialInfo);

          // validate and finalize the signup
          this.httpClient.post('https://fido.workshop/api/v1/member/register/device/finish', registerResponse, {
            headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${this.appStateService.getRawToken()}`,
            },
          }).pipe(
              catchError((err: HttpErrorResponse) => {
                console.error('Signup of new key failed (finish call)', err);
                result$.next(err.status === HttpStatusCode.Unauthorized ? AUTH_STATE.NOT_VERIFIED : AUTH_STATE.HTTP_ERROR);
                result$.complete();
                return throwError(err);
              })
            ).subscribe(() => {
            result$.next(AUTH_STATE.VERIFIED);
            result$.complete();
          });
        }).catch(err => {
          console.log(err);
          result$.next(AUTH_STATE.AUTHENTICATOR_ERROR);
          result$.complete();
        });
      },
      error: (err: HttpErrorResponse) => {
        console.error(`Error while attempting new key registration'`, err);
        result$.next(err.status === HttpStatusCode.BadRequest ? AUTH_STATE.NOT_VERIFIED : AUTH_STATE.HTTP_ERROR);
        result$.complete();
      }
    });

    return result$.asObservable();
  }
}
