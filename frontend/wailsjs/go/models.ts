export namespace models {
	
	export class Product {
	    id: number;
	    name: string;
	    processingTime: number;
	    timeCalculation: string;
	
	    static createFrom(source: any = {}) {
	        return new Product(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.processingTime = source["processingTime"];
	        this.timeCalculation = source["timeCalculation"];
	    }
	}

}

