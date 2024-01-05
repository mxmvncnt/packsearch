use actix_web::{get, HttpResponse, Responder, web};
use actix_web::web::Data;
use serde::Serialize;
use sqlx::FromRow;
use crate::database_structs::{AppState};

#[derive(Serialize, FromRow)]
struct Package {
    id: i64,
    human_name: String,
    name: String,
    latest_version: String,
    description: String,
    keywords: Vec<String>,
    homepage: String,
    developer: Vec<String>,
}

#[get("/search/{query}")]
async fn search_service(state: Data<AppState>, info : web::Path<String>) -> impl Responder {
    let info = info.into_inner();

    println!("{}", info);

    let result = sqlx::query_as::<_, Package>(
        "
            SELECT
                id,
                human_name,
                name,
                latest_version,
                description,
                keywords,
                homepage,
                developer
            FROM package;
            ")
        .fetch_all(&state.db)
        .await;

    match result {
        Ok(packages) => HttpResponse::Ok().json(packages),
        Err(error) => {
            println!("{}", error.to_string());
            HttpResponse::NotFound().json("No packages found.")
        }
    }
}