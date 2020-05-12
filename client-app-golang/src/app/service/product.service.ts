import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Product } from "../entities/product.entity";

@Injectable()
export class ProductService {
  private BASE_URL: string = "http://localhost:6535/api/product/";

  constructor(private httpClient: HttpClient) {}
  findAll() {
    return this.httpClient
      .get(this.BASE_URL + "findall")
      .toPromise()
      .then((res) => res as Product);
  }
  find(id: string) {
    return this.httpClient
      .get(this.BASE_URL + "find/" + id)
      .toPromise()
      .then((res) => res as Product);
  }
  create(product: Product) {
    return this.httpClient
      .post(this.BASE_URL + "create", product)
      .toPromise()
      .then((res) => res as Product);
  }
  delete(id: string) {
    return this.httpClient
      .delete(this.BASE_URL + "delete/" + id)
      .toPromise()
      .then((res) => res as Product);
  }
  update(product: Product) {
    return this.httpClient
      .put(this.BASE_URL + "update", product)
      .toPromise()
      .then((res) => res as Product);
  }
}
