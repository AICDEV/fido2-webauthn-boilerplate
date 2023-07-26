import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { LoginService } from '../../service/login.service';
import { AUTH_STATE } from '../../model/types';
import { SnackbarService } from '../../service/snackbar.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.view.component.html',
  styleUrls: ['./login.view.component.scss'],
})
export class LoginViewComponent {
  public loginForm: FormGroup;

  constructor(
    private router: Router,
    private formBuilder: FormBuilder,
    private loginService: LoginService,
    private snackbarService: SnackbarService,
  ) {
    this.loginForm = this.formBuilder.group({
      email: new FormControl('', [Validators.required, Validators.email]),
    });

    // TODO: add auto-login with automatic redirect
  }

  public login(): void {
    if (!window.PublicKeyCredential) {
      this.snackbarService.showSnackbarWithMessage('Login failed because of missing Browser-Support');
      return;
    }

    this.loginService.login(this.loginForm.get('email')?.value).subscribe((loginState) => {
      this.snackbarService.showSnackbarForAuthState(loginState);

      if (loginState === AUTH_STATE.VERIFIED) {
        this.router.navigate(['/main']);
      }
    });
  }
}
