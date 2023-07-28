import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of, Subject } from 'rxjs';
import {
  AUTH_STATE,
  ILoginData,
} from '../model/types';
import { SnackbarService } from './snackbar.service';
import { AppStateService } from './appstate.service';


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
    const resultSub = new Subject<AUTH_STATE>();

    // call 'https://fido.workshop/api/v1/service/authenticate/begin' with post request
    // content type for response: WrappedPublicKeyCredentialRequestOptionsJSON
    /* post body: {
      email: email // the email property given to the login fn
    }
    */

    // handle WrappedPublicKeyCredentialRequestOptionsJSON response
    // map into CredentialRequestOptions
    // call navigator.credentials.get(options) with these CredentialRequestOptions
    // create ILoginData from navigator.credentials.get result
    // call finalizeLogin fn with loginData
    // trigger resultSub.next with correct AUTH_STATE and complete the observable stream

    // login fn returns overall AUTH_STATE observable depending on success/failure
    return resultSub.asObservable();
  }

  private finalizeLogin(loginData: ILoginData): Observable<boolean> {
    // POST call to 'https://fido.workshop/api/v1/service/authenticate/finish'
    // body content is the given ILoginData
    // response object is of type IFinalizeLoginResponse
    // decode JWT from response
    // call this.appStateService.setUser with user metadata from token
    // store current JWT into app state via call to this.appStateService.setToken
    // set login state to true by calling this.appStateService.setLoginState
    // in error case call this.snackbarService.showSnackbarWithMessage('Error while decoding the jwt'); and return false
    // return true if login finalization were successful

    return of(false);
  }
}
