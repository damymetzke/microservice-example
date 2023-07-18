use actix_web::{get, web, App, HttpServer, Responder, HttpRequest};

#[get("/users")]
async fn list_users() -> impl Responder {
    "Hello users!"
}

#[get("/users/{id}")]
async fn get_user(id: web::Path<i32>) -> impl Responder {
    format!("Hello user {id}!")
}

#[get("*")]
async fn failed(req: HttpRequest) -> impl Responder {
    println!("{}", req.path());
    ""
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(list_users).service(get_user))
        .bind(("0.0.0.0", 80))?
        .run()
        .await
}
