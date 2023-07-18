use actix_web::{get, web, App, HttpServer, Responder};

#[get("/users")]
async fn list_users() -> impl Responder {
    "Hello users!"
}

#[get("/users/{id}")]
async fn get_user(id: web::Path<i32>) -> impl Responder {
    format!("Hello user {id}!")
}

#[actix_web::main] // or #[tokio::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(list_users).service(get_user))
        .bind(("127.0.0.1", 8080))?
        .run()
        .await
}
