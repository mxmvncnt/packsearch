mod database_structs;
mod package;
mod packages_list;
mod search;
mod variation;

use crate::database_structs::AppState;
use crate::package::{get_package_service, post_package_service};
use crate::packages_list::packages;
use crate::search::search_service;
use crate::variation::get_variations_service;
use actix_web::{web::Data, App, HttpServer, Responder};
use dotenv::dotenv;
use sqlx::postgres::PgPoolOptions;

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

    HttpServer::new(move || {
        App::new()
            .app_data(Data::new(AppState { db: pool.clone() }))
            .service(packages)
            .service(search_service)
            .service(get_variations_service)
            .service(get_package_service)
            .service(post_package_service)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
