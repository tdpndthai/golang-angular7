import { Component, OnInit } from "@angular/core";
import { Product } from "./entities/product.entity";
import { ProductService } from "./service/product.service";
import { FormGroup, FormBuilder } from "@angular/forms";

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"],
})
export class AppComponent implements OnInit {
  products: Product[] = []; //products: Product[] - khai báo biến,còn cái này là khởi tạo
  productId: string = "";
  product: Product = {}; //khởi tạo để nó hiểu là object
  productForm: FormGroup;
  updateProductForm: FormGroup;

  constructor(
    private productService: ProductService,
    private formBuilder: FormBuilder
  ) {}
  title = "client-app-golang";
  ngOnInit() {
    this.productService.findAll().then(
      (res: any) => {
        //chuyển giá trị cho rest là any
        //this.products.push(res);
        //console.log(typeof this.products);
        this.products = res;
      },
      (error) => {
        console.log(error);
      }
    );
    this.productForm = this.formBuilder.group({
      name: "",
      price: 0,
      status: false,
    });
    this.updateProductForm = this.formBuilder.group({
      id: "",
      name: "",
      price: 0,
      status: false,
    });
  }

  search() {
    //console.log(this.productId)
    this.productService.find(this.productId).then(
      (res) => {
        this.product = res;
      },
      (error) => {
        console.log(error);
      }
    );
  }
  save() {
    var product: Product = this.productForm.value;
    //console.log(product)
    this.productService.create(product).then(
      (res) => {
        this.product = res;
      },
      (error) => {
        console.log(error);
      }
    );
  }

  delete() {
    //console.log(this.productId)
    this.productService.delete(this.productId).then(
      (res) => {
        this.product = res;
      },
      (error) => {
        console.log(error);
      }
    );
  }

  update() {
    var product: Product = this.updateProductForm.value;    
    //console.log(product)
    this.productService.update(product).then(
      (res) => {
        this.product = res;
      },
      (error) => {
        console.log(error);
      }
    );
  }
}
