import { ResourceBean } from '../resource-bean';

export class DataSourceBean extends ResourceBean {
    public adress: string;
    public pipes: string;
    public body: string;
    public resultset: any[];
}
