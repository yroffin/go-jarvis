import { ResourceBean } from './resource-bean';
import { CronBean } from './cron-bean';

export class TriggerBean extends ResourceBean {
    public crons: CronBean[];
    public devices: any;
    public processes: any;
}
