import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '../../../core';

@Component({
  selector: 'app-users-admin-preview',
  styleUrls: ['./users-admin.component.css'],
  templateUrl: './users-admin-preview.html'
})
export class UsersAdminPreviewComponent {
  constructor (
    private userService: UserService,
    private router: Router
  ) {}

  @Input() user: User;

  deleteUser(event) {
    event.stopPropagation()

    this.userService.destroy(this.user.id)
      .subscribe(
        success => {
          this.router.navigateByUrl('.', { skipLocationChange: true })
          .then(() => {
            this.router.navigate(['/']);
        }); 
        }
      );
  }

  modifySong(event, id) {
    event.stopPropagation()
    this.router.navigateByUrl("/editor/" + id);
  }
}
