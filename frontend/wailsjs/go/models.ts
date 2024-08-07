export namespace backend {
	
	export class Event {
	    id: number;
	    title: string;
	    description: string;
	    date: string;
	
	    static createFrom(source: any = {}) {
	        return new Event(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.date = source["date"];
	    }
	}

}

