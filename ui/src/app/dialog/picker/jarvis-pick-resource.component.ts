/* 
 * Copyright 2016 Yannick Roffin.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Component, Inject, AfterViewInit } from "@angular/core";
import { MatDialogRef, MAT_DIALOG_DATA, MatTableDataSource, MatPaginator } from "@angular/material";
import { ResourceBean } from "../../model/resource-bean";
import { ViewChild } from "@angular/core";
import { NotifyCallback } from "../../class/jarvis-resource";
import { PickerBean } from "../../model/picker-bean";

export interface DialogData {
  target: NotifyCallback<ResourceBean>,
  title: string,
  resources: ResourceBean[],
  picker: PickerBean,
}

@Component({
  selector: 'dialog-pick-resource',
  templateUrl: 'jarvis-pick-resource.component.html',
  styleUrls: ['jarvis-pick-resource.component.css']
})
export class DialogPickResource implements AfterViewInit {

  public myMatResources = new MatTableDataSource([])
  @ViewChild('paginator') paginator: MatPaginator;

  constructor(
    public dialogRef: MatDialogRef<DialogPickResource>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData) {
    this.myMatResources.data = data.resources;
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

  /**
   * validate and close
   */
  public validate(resource: ResourceBean) {
    if (resource) {
      this.data.target.notify(this.data.picker, resource);
    }
    this.dialogRef.close();
  }

  applyFilter(filterValue: string) {
    filterValue = filterValue.trim(); // Remove whitespace
    filterValue = filterValue.toLowerCase(); // MatTableDataSource defaults to lowercase matches
    this.myMatResources.filter = filterValue;
  }
  
  /**
   * Set the paginator after the view init since this component will
   * be able to query its view for the initialized paginator.
   */
  ngAfterViewInit() {
    this.myMatResources.paginator = this.paginator;
  }
}