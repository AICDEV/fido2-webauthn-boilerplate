import { Component } from '@angular/core';
import { AdminService } from '../../service/admin.service';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.view.component.html',
  styleUrls: ['./admin.view.component.scss'],
})
export class AdminViewComponent {
  public protectedDataFromServer: { name: string } | undefined;

  constructor(
    private adminService: AdminService,
  ) {
  }

  public fetchData(): void {
    this.adminService.fetchProtectedData().subscribe(res => {
      this.protectedDataFromServer = res;
    });
  }
}
