use crate::database_structs::{AppState, Package};
use actix_web::web::Data;
use actix_web::{get, HttpResponse, Responder};

#[get("/packages")]
async fn packages(state: Data<AppState>) -> impl Responder {
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
            ",
    )
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
