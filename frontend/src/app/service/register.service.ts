import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { lastValueFrom, Observable, of, Subject } from 'rxjs';
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

    // POST call to 'https://fido.workshop/api/v1/service/signup/begin'
    /* body:
    * {
      displayName: `${firstName} ${lastName}`,
      firstName: firstName,
      lastName: lastName,
      email: mail,
    }
    * */
    // return type of response IRegisterResponse; @see app/model/types.ts

    // map IRegisterResponse into options for navigator.credentials.create call
    // perform navigator.credentials.create call with these options
    // create registerResponse (type RegistrationResponseJSON) from navigator.credentials.create fn call response
    // do POST call to 'https://fido.workshop/api/v1/service/signup/finish' with registerResponse as body
    // trigger result$.next according to success or error with AUTH_STATE and complete the stream

    return result$.asObservable();
  }

  private createFinalRegistrationValidationData(credentialInfo: RegistrationCredential): RegistrationResponseJSON  {
    // map credentialInfo parameter into RegistrationResponseJSON; @see app/model/types.ts

    // replace me... placeholder so it compiles
    return { } as any as RegistrationResponseJSON;
  }

  private credentialCreateWithRegisterResponse(registerResponse: IRegisterResponse): Promise<RegistrationCredential | null> {
    // map registerResponse parameter to navigator.credentials.create options call it and return the Promise to the result

    // replace me... placeholder so it compiles
    return lastValueFrom(of(null));
  }

  public registerNewKey(): Observable<AUTH_STATE> {
    const result$ = new Subject<AUTH_STATE>();

    // POST call to 'https://fido.workshop/api/v1/member/register/device/begin'
    // body is undefined
    // headers with bearer token which you can get by calling this.appStateService.getRawToken()
    // response type IRegisterResponse; @see app/model/types.ts

    // call credentialCreateWithRegisterResponse fn to map registerResponse to navigator.credentials.create options call it and return the Promise to the result
    // map result into registerResponse of type RegistrationResponseJSON
    // POST call to 'https://fido.workshop/api/v1/member/register/device/finish' with registerResponse as body and headers containing the bearer token
    // trigger result$.next with according AUTH_STATE and close the observable response stream

    return result$.asObservable();
  }
}
