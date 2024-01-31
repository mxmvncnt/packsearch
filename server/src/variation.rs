use crate::database_structs::{AppState, FullVariation};
use actix_web::web::Data;
use actix_web::{get, web, HttpResponse, Responder};

#[get("/package/{package_id}/variations")]
async fn get_variations_service(
    state: Data<AppState>,
    package_id: web::Path<String>,
) -> impl Responder {
    let package_id = package_id.into_inner().parse::<i64>().unwrap();

    let variations_result = sqlx::query_as::<_, FullVariation>(
        "
            SELECT variation.id      AS id,
                   package_id,
                   distro_id,
                   variation.name    AS name,
                   variation.version AS version,
                   package_url,
                   download_url,
                   d.name            AS distro_name,
                   d.version         AS distro_version
            FROM variation
                     INNER JOIN distro d ON d.id = variation.distro_id
            WHERE package_id = $1;
            ",
    )
    .bind(package_id)
    .fetch_all(&state.db)
    .await;

    match variations_result {
        Ok(variations) => {
            if variations.len() == 0 {
                HttpResponse::NotFound().json("No variations found for this package.")
            } else {
                HttpResponse::Ok().json(variations)
            }
        }
        Err(error) => {
            println!("{}", error.to_string());
            HttpResponse::NotFound().json("An error has occurred.")
        }
    }
}
