import { CommonModule } from "@angular/common";
import { Component, Inject } from "@angular/core";
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from "@angular/forms";
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from "@angular/material/dialog";
import { AddReminder } from "src/app/interfaces/add-reminder";
import { Reminder } from "src/app/interfaces/reminder";
import { MaterialModule } from "src/app/material.module";
import { ReminderService } from "src/app/services/reminderService/reminder.service";
import Swal from "sweetalert2";

@Component({
    selector: 'add-reminder',
    templateUrl : './add-reminder.modal.html',
    standalone: true,
    imports: [MaterialModule, ReactiveFormsModule, FormsModule, CommonModule]
})
export class AddReminderModal{
    constructor(public dialogRef : MatDialogRef<AddReminderModal>,
        private _service : ReminderService
    ){ }

    loadingRequest: boolean = false;


    reminderForm = new FormGroup({
        'title' : new FormControl('', [Validators.required]),
        'description' : new FormControl('', [Validators.required]),
    })

    addReminder(){
        const newReminder : AddReminder = {
            description : this.reminderForm.get('description')?.value?.toString(),
            title : this.reminderForm.get('title')?.value?.toString()
        }
        this.loadingRequest = true;
        this._service.AddReminders(newReminder).subscribe((response)=>{
            this.loadingRequest = false;
            Swal.fire('', 'Reminder added successfully', 'success').then(()=>{this.dialogRef.close()});
        }, (error)=>{
            this.loadingRequest = false;
            Swal.fire('Oops', 'Error while adding that reminder!', 'error');
        })
    }
     
}

@Component({
    selector: 'update-reminder',
    templateUrl : './add-reminder.modal.html',
    standalone: true,
    imports: [MaterialModule, ReactiveFormsModule, FormsModule, CommonModule]
})
export class UpdateReminderModal{
    constructor(public dialogRef : MatDialogRef<AddReminderModal>,
        @Inject(MAT_DIALOG_DATA) public reminder: Reminder,
        private _service : ReminderService
    ){ }

    loadingRequest: boolean = false;


    reminderForm = new FormGroup({
        'title' : new FormControl(this.reminder.title),
        'description' : new FormControl(this.reminder.description),
    })

    addReminder(){
        const newReminder : AddReminder = {
            description : this.reminderForm.get('description')?.value?.toString(),
            title : this.reminderForm.get('title')?.value?.toString(),
            reminderId : this.reminder.reminderId

        }
        this.loadingRequest = true;
        this._service.Updatereminder(newReminder).subscribe((response)=>{
            this.loadingRequest = false;
            Swal.fire('', 'Reminder updated successfully', 'success').then(()=>{this.dialogRef.close()});
        }, (error)=>{
            this.loadingRequest = false;
            Swal.fire('Oops', 'Error while updating that reminder!', 'error');
        })
    }
     
}