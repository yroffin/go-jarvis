/**
 * default node model
 */
export class NodeBean {
    public id: string;
    public label: string;
    public title: string;
    public group: any;
}

/**
 * default edge model
 */
export class EdgeBean {
    public from: string;
    public to: string;
    public label: string;
    public title: string;
    public smooth: boolean;
    public arrows: any;
    public data: any;
}

/**
 * default graph model
 */
export class GraphBean {
    public nodes: NodeBean[];
    public edges: EdgeBean[];
    public options: any;
}
