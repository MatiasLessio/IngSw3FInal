import { Component } from '@angular/core';
import { GlobalConstants } from '../utilities/globalConstants';
import { ReminderService } from '../services/reminderService/reminder.service';
import { Reminder } from '../interfaces/reminder';
import { CommonModule } from '@angular/common';
import { MaterialModule } from '../material.module';
import { MatDialog } from '@angular/material/dialog';
import { AddReminderModal, UpdateReminderModal } from './modals/crud-reminders.component';
import { Title } from '@angular/platform-browser';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, MaterialModule],
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent {
  
  user:string= '';
  ngOnInit():void{
    if(!sessionStorage.getItem('token')){
      window.location.href = 'Login'
    }
    this.user = sessionStorage.getItem('username')!;
    this.getReminders();
  }

  constructor(private _reminderService : ReminderService,
    public dialog: MatDialog
  ) {

  }
  reminders : Reminder[] = [];
  getReminders(){
    this._reminderService.GetReminders().subscribe((response)=>{
      this.reminders = response;
    })
  }

  addReminder(){
    const dialogRef = this.dialog.open(AddReminderModal, {
      width: 'auto',
      disableClose : true
    })
    dialogRef.afterClosed().subscribe(()=>{
      this.getReminders();
    })
  }
  editReminder(reminder : Reminder){
    const dialogRef = this.dialog.open(UpdateReminderModal, {
      width: 'auto',
      disableClose : true,
      data: reminder
    })
    dialogRef.afterClosed().subscribe(()=>{
      this.getReminders();
    })
  }
  deleteReminder(reminder: Reminder) {
    Swal.fire({
        title: 'Want to delete ' + reminder.title + ' ?',
        icon: 'warning',
        showCancelButton: true,
        focusConfirm: false,
        confirmButtonText: 'Yes',
        cancelButtonText: 'No',
        reverseButtons: true
    }).then((result) => {
        if (result.isConfirmed) {
            this._reminderService.DeleteReminders(reminder).subscribe((response) => {
                    Swal.fire('', 'Reminder deleted successfully!', 'success').then(() => { this.getReminders() });
            }, (error) => {
                Swal.fire('Oops', 'Error while deleting reminder!', 'error');
            })
        }
    });
}
}
