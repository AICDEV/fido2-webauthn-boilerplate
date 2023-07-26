import { NgModule } from '@angular/core';
import { MainViewComponent } from './main.view.component';
import { MainViewRoutingModule } from './main.view-routing.module';


@NgModule({
  declarations: [
    MainViewComponent,
  ],
  imports: [
    MainViewRoutingModule,
  ],
  providers: [],
})
export class MainViewModule { }
