use actix_web::{get, web, App, HttpRequest, HttpServer, Responder};
use serde::Serialize;

#[derive(Serialize)]
struct User {
    id: i32,
    name: &'static str,
    email: &'static str,
}

const USERS: [User; 3] = [
    User {
        id: 0,
        name: "Alice",
        email: "alice@example.com",
    },
    User {
        id: 1,
        name: "Bob",
        email: "bob@example.com",
    },
    User {
        id: 2,
        name: "Charly",
        email: "charly@example.com",
    },
];

#[get("/users")]
async fn list_users() -> impl Responder {
    serde_json::to_string(&USERS)
}

#[get("/users/{id}")]
async fn get_user(id: web::Path<i32>) -> impl Responder {
    match USERS.iter().find(|user| user.id == *id) {
        Some(user) => serde_json::to_string(user),
        None => Ok(String::new()),
    }
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
