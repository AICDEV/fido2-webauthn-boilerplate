import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RegisterViewComponent } from './register.view.component';
import { RegisterViewRoutingModule } from './register.view-routing.module';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatInputModule } from '@angular/material/input';
import { ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatSnackBarModule } from '@angular/material/snack-bar';


@NgModule({
  declarations: [
    RegisterViewComponent,
  ],
  imports: [
    CommonModule,
    RegisterViewRoutingModule,
    MatFormFieldModule,
    MatSelectModule,
    MatInputModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatCardModule,
    MatSnackBarModule,
  ],
  providers: [],
})
export class RegisterViewModule { }
