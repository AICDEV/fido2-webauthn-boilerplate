import { Injectable } from '@angular/core';
import { AUTH_STATE } from '../model/types';
import { MatSnackBar } from '@angular/material/snack-bar';


@Injectable({
  providedIn: 'root'
})
export class SnackbarService {

  constructor(
    private matSnackBar: MatSnackBar,
  ) {}

  public showSnackbarForAuthState(authState: AUTH_STATE): void {
    switch (authState) {
      case AUTH_STATE.VERIFIED:
        this.matSnackBar.open('You are logged in', undefined, {
          duration: 2000,
        });
        break;
      case AUTH_STATE.NOT_VERIFIED:
        this.matSnackBar.open('Login failed, because verification failed.', undefined, {
          duration: 2000,
        });
        break;
      case AUTH_STATE.AUTHENTICATOR_ERROR:
        this.matSnackBar.open('Login failed. Authentication error.', undefined, {
          duration: 2000,
        });
        break;
      case AUTH_STATE.UNKNOWN_USER:
        this.matSnackBar.open('Login failed, because user is unknown.', undefined, {
          duration: 2000,
        });
        break;
      case AUTH_STATE.HTTP_ERROR:
        this.matSnackBar.open('Login failed. Network issue.', undefined, {
          duration: 2000,
        });
        break;
    }
  }

  public showSnackbarWithMessage(message: string): void {
    this.matSnackBar.open(message, undefined, {
      duration: 2000,
    });
  }
}
