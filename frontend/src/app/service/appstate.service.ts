import { Injectable } from '@angular/core';
import { IUser } from '../model/user.type';
import { BehaviorSubject, map, Observable } from 'rxjs';
import { JwtPayload } from 'jwt-decode';



export interface IAppState  {
  loggedIn: boolean;
  user: IUser | null;
  token: JwtPayload | null;
  rawToken: string | null;
}

@Injectable({
  providedIn: 'root',
})
export class AppStateService {
  private appState$: BehaviorSubject<IAppState>;

  constructor() {
    this.appState$ = new BehaviorSubject<IAppState>({
      user: null,
      loggedIn: false,
      token: null,
      rawToken: null,
    });
  }

  public getAppState$(): Observable<IAppState> {
    return this.appState$.asObservable();
  }

  public setUser(user: IUser): void {
    this.appState$.next({
      ...this.appState$.getValue(),
      user,
    });
  }

  public removeUser(): void {
    this.appState$.next({
      ...this.appState$.getValue(),
      user: null,
    });
  }

  public setLoginState(loggedIn: boolean): void  {
    this.appState$.next({
      ...this.appState$.getValue(),
      loggedIn,
    });
  }

  public logout(): void {
    this.appState$.next({
      token: null,
      rawToken: null,
      user: null,
      loggedIn: false,
    });
  }

  public isLoggedIn(): boolean {
    return this.appState$.getValue().loggedIn;
  }

  public setToken(rawToken: string, token: JwtPayload): void {
    this.appState$.next({
      ...this.appState$.getValue(),
      token: token,
      rawToken: rawToken,
    });
  }

  public getToken$(): Observable<JwtPayload | null> {
    return this.appState$.pipe(map(state => state.token));
  }

  public getToken(): unknown {
    return this.appState$.getValue().token;
  }

  public getUser$(): Observable<IUser | null> {
    return this.appState$.pipe(map(state => state.user));
  }

  public getIsLoggedIn$(): Observable<boolean> {
    return this.appState$.pipe(map(state => state.loggedIn));
  }

  public getRawToken(): string | null {
    return this.appState$.getValue().rawToken;
  }
}
