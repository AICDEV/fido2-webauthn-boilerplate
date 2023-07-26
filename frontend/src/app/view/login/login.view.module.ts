import { NgModule } from '@angular/core';
import { LoginViewComponent } from './login.view.component';
import { LoginViewRoutingModule } from './login.view-routing.module';
import { MatCardModule } from '@angular/material/card';
import { ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatButtonModule } from '@angular/material/button';


@NgModule({
  declarations: [
    LoginViewComponent,
  ],
  imports: [
    LoginViewRoutingModule,
    MatCardModule,
    MatButtonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatSnackBarModule,
  ],
  providers: [],
})
export class LoginViewModule { }
