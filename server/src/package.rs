use crate::database_structs::{AppState, Package};
use actix_web::web::Data;
use actix_web::{get, web, HttpResponse, Responder, post};
use serde::{Deserialize, Serialize};
use sqlx::FromRow;

#[derive(Serialize, Deserialize, FromRow)]
struct UploadPackage {
    human_name: String,
    name: Option<String>,
    latest_version: Option<String>,
    description: Option<String>,
    keywords: Option<Vec<String>>,
    homepage: Option<String>,
    developer: Vec<String>,
}

#[get("/package/{package_id}")]
async fn get_package_service(
    state: Data<AppState>,
    package_id: web::Path<String>,
) -> impl Responder {
    let package_id = package_id.into_inner().parse::<i64>().unwrap();

    println!("{}", package_id);

    let package_result = sqlx::query_as::<_, Package>(
        "
            SELECT *
            FROM package
            WHERE id = $1;
            ",
    )
    .bind(package_id)
    .fetch_all(&state.db)
    .await;

    match package_result {
        Ok(package) => HttpResponse::Ok().json(package),
        Err(error) => {
            println!("{}", error.to_string());
            HttpResponse::NotFound().json("An error has occurred.")
        }
    }
}

#[post("/package/new")]
async fn post_package_service(
    state: Data<AppState>,
    package: web::Json<UploadPackage>
) -> impl Responder {
    println!("{:?}", package.name);
    return HttpResponse::Ok();
}
