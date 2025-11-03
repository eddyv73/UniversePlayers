mod httpconnector;

fn main() {
    let url = "https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/teams?limit=32";
    httpconnector::request_message(url);
    httpconnector::make_request(url);
}

