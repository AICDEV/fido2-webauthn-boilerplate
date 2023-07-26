import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { RegisterService } from '../../service/register.service';
import { AppStateService } from '../../service/appstate.service';
import { Router } from '@angular/router';
import { AUTH_STATE } from '../../model/types';
import { LoginService } from '../../service/login.service';
import { SnackbarService } from '../../service/snackbar.service';


@Component({
  selector: 'app-register',
  templateUrl: './register.view.component.html',
  styleUrls: ['./register.view.component.scss'],
})
export class RegisterViewComponent {
  public registerForm: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private registerService: RegisterService,
    private loginService: LoginService,
    private appStateService: AppStateService,
    private router: Router,
    private snackbarService: SnackbarService,
  ) {
    this.registerForm = this.formBuilder.group({
      email: new FormControl('', [Validators.required, Validators.email]),
      firstName: new FormControl('', [Validators.required]),
      lastName: new FormControl('', [Validators.required]),
    });
  }

  public register(): void {
    this.registerService.register(
      this.registerForm.get('email')?.value,
      this.registerForm.get('firstName')?.value,
      this.registerForm.get('lastName')?.value
    ).subscribe((registerAuthState: AUTH_STATE) => {
      if (registerAuthState === AUTH_STATE.VERIFIED) {
        this.snackbarService.showSnackbarWithMessage('Account created, now login with your key.');
        this.loginService.login(this.registerForm.get('email')?.value).subscribe(loginState => {

          this.snackbarService.showSnackbarForAuthState(loginState);

          if (loginState === AUTH_STATE.VERIFIED) {
            this.router.navigate(['/main']);
          }
        });
      } else {
        this.snackbarService.showSnackbarForAuthState(registerAuthState);
      }
    });
  }
}
