use crate::database_structs::{AppState, Package, Variation};
use actix_web::web::Data;
use actix_web::{get, web, HttpResponse, Responder};

struct Response {
    id: i64,
    human_name: String,
    name: String,
    latest_version: String,
    description: String,
    keywords: Vec<String>,
    homepage: String,
    developer: Vec<String>,
    variations: Vec<Variation>,
}

#[get("/search/{query}")]
async fn search_service(state: Data<AppState>, query: web::Path<String>) -> impl Responder {
    let query = query.into_inner().to_lowercase();

    let potential_packages = sqlx::query_as::<_, Package>(
        "
            SELECT DISTINCT package.id,
                            human_name,
                            package.name,
                            latest_version,
                            description,
                            keywords,
                            homepage,
                            developer
            FROM package
                     INNER JOIN public.variation v ON package.id = v.package_id
            WHERE dmetaphone($1) ILIKE dmetaphone(human_name)
               OR dmetaphone($1) = ANY (ARRAY(SELECT dmetaphone(element) FROM unnest(keywords) AS element))
               OR dmetaphone($1) = ANY (ARRAY(SELECT dmetaphone(element) FROM unnest(developer) AS element))
               OR dmetaphone($1) ILIKE dmetaphone(human_name)
               OR dmetaphone($1) ILIKE dmetaphone(package.name)
               OR dmetaphone($1) ILIKE dmetaphone(description)
               OR dmetaphone($1) ILIKE dmetaphone(v.name);
            ",
    )
    .bind(query)
    .fetch_all(&state.db)
    .await;

    match potential_packages {
        Ok(packages) => {
            if packages.len() == 0 {
                HttpResponse::NotFound().json("No packages found for this query.")
            } else {
                HttpResponse::Ok().json(packages)
            }
        }
        Err(error) => {
            println!("{}", error.to_string());
            HttpResponse::NotFound().json("An error has occured.")
        }
    }
}
