use actix_web::{get, HttpResponse, Responder, web};
use actix_web::web::{Data, resource};
use serde::Serialize;
use sqlx::{FromRow, query};
use crate::database_structs::{AppState, Variation};

#[get("/package/{package_id}/variations")]
async fn get_variations_service(state: Data<AppState>, package_id : web::Path<String>) -> impl Responder {
    let package_id = package_id.into_inner().parse::<i64>().unwrap();

    println!("{}", package_id);

    let variations_result = sqlx::query_as::<_, Variation>(
        "
            SELECT *
            FROM variation
            WHERE package_id = $1;
            ")
        .bind(package_id)
        .fetch_all(&state.db)
        .await;

    match variations_result {
        Ok(variations) => {
            if (variations.len() == 0) {
                HttpResponse::NotFound().json("No variations found for this package.")
            } else {
                HttpResponse::Ok().json(variations)
            }
        },
        Err(error) => {
            println!("{}", error.to_string());
            HttpResponse::NotFound().json("An error has occured.")
        }
    }
}