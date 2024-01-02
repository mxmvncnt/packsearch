use actix_web::{App, get, HttpResponse, HttpServer, Responder};
use sqlx::{postgres::PgPoolOptions, Pool, Postgres};
use dotenv::dotenv;

#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world!")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();

    // base connection code from https://github.com/bocksdin/rust-sqlx/tree/main
    let database_url = std::env::var("DB_URL").expect("DB_URL must be set");
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&database_url)
        .await
        .expect("Error building a connection pool");

    HttpServer::new(|| App::new()
        .service(hello)
    )
        .bind(("127.0.0.1", 8080))?
        .run()
        .await
}
