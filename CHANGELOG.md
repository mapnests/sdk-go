# Changelog — `multi-stop-point` branch

Scope: all changes on this branch since it diverged from `main` (merge-base
`59592d9` — "updated error message format"), covering both the committed
commit `0197ae7` ("Added new func and validation") and the current
uncommitted working-tree changes.

## New APIs

### Geocode (`gecode.go`)
- New `Client.Geocode(ctx, GeocodeRequest)` method — `GET /geomap/api/v2/geocode`.
- `GeocodeRequest{Query, Limit, AcceptedLanguage, XRequestID}`.
- `GeocodeResponse` returns `[]GeocodeData` (place_id, lat/lon, category, type,
  place_rank, importance, addresstype, name, display_name, address).
- Validates `XRequestID`, `AcceptedLanguage`, `Limit` (max `GeocodeMaxLimit = 10`),
  and normalizes the query string.
- Note: the file is named `gecode.go` (typo for "geocode") — the type/function
  names themselves are spelled correctly.

### Multi-source route summary without geometry (`multi_source_summary_without_geometry.go`)
- New `Client.MultiSourceRouteSummaryWithoutGeometry(ctx, request)` method —
  `POST /routemap/api/v1/routes/multi-source-summary-without-geometry`.
- Returns `RouteSummaryWithoutGeometry{ID, DistanceMeters, EtaSeconds, Reachable}`
  per source.
- Request type churn: introduced as its own
  `MultiSourceRouteSummaryRequestWithoutGeometry` struct in the commit, then in
  the uncommitted changes switched to reuse the existing
  `MultiSourceRouteSummaryRequest` type instead (the now-unused struct is left
  commented out in the file rather than removed).
- Validates `XRequestID` and lat/lon for every source plus the destination.

### ETA APIs (new, uncommitted)
- `eta_without_geometry_by_mode.go` — `Client.EtaWithoutGeometryByMode(ctx, ETAByModeRequest)`,
  `GET /eta/api/v1/get-eta`. Returns `[]EtaWithoutGeometryByModeData{From, To, Distance, Eta}`.
- `eta_with_geometry_by_mode.go` — `Client.EtaWithGeometryByMode(ctx, ETAByModeRequest)`,
  `GET /eta/api/v1/get-eta/details`. Returns a single `RouteData` payload.
- `eta_with_geometry.go` — `Client.EtaWithGeometry(ctx, ETAsRequest)`,
  `GET /eta/api/v1/get-eta/details/multiple-modes`. Returns per-mode results
  for car/CNG/motorcycle in one call (`ETAsModesData`).
- `ETAByModeRequest{FromLat, FromLon, ToLat, ToLon, Mode, NoOfRoutes, XRequestID}`;
  `NoOfRoutes` is validated against a new `ETAMaxRoutes = 5` limit.
- `validate_eta_mode.go` — new `ValidateEtaMode(mode Mode)`, restricting ETA
  requests to `car`, `cng`, and `motorcycle` (`ErrInvalidEtaMode`).
- All three ETA methods validate from/to lat-lon and `XRequestID`; the
  by-mode variants additionally validate `Mode` and `NoOfRoutes`.
- `client.go` and `secure_request.go`'s `HTTPMethodMap` updated to register
  the three new labels (`etaWithoutGeometryByMode`, `etaWithGeometryByMode`,
  `etaWithGeometry`) as GET requests.

## New request-level validation

### `validate_xrequestid.go` (new)
- `ValidateXRequestID(*string)` enforces the W3C `traceparent` format
  (`<2 hex version>-<32 hex trace-id>-<16 hex parent-id>-<2 hex flags>`) and
  rejects the reserved-invalid version `ff` and all-zero trace/parent IDs.
- Wired into every request method: `Autocomplete`, `DetailsByPlaceID`,
  `DistanceMatrix`, `DistanceMatrixDetails`, `MultiSourceRouteSummary`,
  `MultiSourceRouteSummaryWithoutGeometry`, `MultiStopPoints`,
  `PairWiseRouteSummary`, `Reverse`, `Search`, `SearchByRadius`, `SnapToRoad`,
  `Geocode`, and all three new ETA methods.
- `secure_request.go` forwards the value as an `x-request-id` header via new
  `XRequestIDFromJSON(jsonRequest string) string`, extracted from the
  marshaled request body.

### `validate_accepted_language.go` (new)
- `ValidateAcceptedLanguage(*string)` restricts the field to `"en"` or `"bn"`
  (`ErrInvalidAcceptedLanguage`).
- New `AcceptedLanguage *string` field added to `AutoCompleteRequest`,
  `ReverseRequest`, `SearchRequest`, `SearchByRadiusRequest`, and
  `GeocodeRequest`; threaded through to the corresponding query string in
  `build_url_from_json.go` as `&acceptedLanguage=...`.

### `validate_limit_radius.go` (new)
- `ValidateLimit(limit *int64, max int64)` / `ValidateRadius` /
  `ValidateRadiusPtr` / `ValidateRadiusRequiresLatLon`.
- New per-endpoint limits: `SearchMaxLimit=20`, `SearchMaxRadius=150000`,
  `SearchByRadiusMaxLimit=20`, `SearchByRadiusMaxRadius=150000`,
  `GeocodeMaxLimit=10`, `AutocompleteMaxLimit=20`, `AutocompleteMaxRadius=5000`,
  and (uncommitted) `ETAMaxRoutes=5`.
- Applied to `Search`, `SearchByRadius`, `Autocomplete`, `Geocode`, and the
  ETA-by-mode methods; `ValidateRadiusRequiresLatLon` additionally enforces
  that `Search`/`Autocomplete` requests supplying a `Radius` must also supply
  `Lat`/`Lon`.

## Other endpoint changes

- **`autocomplete.go`**: added `Radius *int64` and `AcceptedLanguage *string`
  to `AutoCompleteRequest`; now validates `XRequestID`, `AcceptedLanguage`,
  `Limit`, `Radius`, and radius-requires-lat/lon before the request.
- **`snap_to_road.go`**: added `NumberOfNearestRoad *int32` to
  `SnapToRoadRequest`; now validates `XRequestID`.
- **`build_url_from_json.go`**:
  - Reorganized into "Geomap APIs" / "Routemap APIs" / "ETA APIs" sections.
  - Split the combined `distanceMatrix`/`distanceMatrixDetails` case into two
    explicit cases.
  - Split the combined `search`/`searchByRadius` case into two explicit cases
    (radius search now hits `/geomap/api/v2/search/radius` instead of the
    previous `/geomap/api/v2/search/radius`-via-shared-path construction —
    same resulting path, clearer code).
  - Added `acceptedLanguage` query param wiring for `reverse`, `geocode`,
    `search`, `searchByRadius`, and `autocomplete`.
  - Added the new `geocode`, `multiSourceRouteSummaryWithoutGeometry`, and the
    three `eta*` cases.
  - **`baseURL` changed from `https://engine.mapnests.com` to
    `http://192.168.169.58:9080`** — looks like a local/debug override left
    in; should be reverted before merging.
- **`secure_request.go`**: added a stray `fmt.Println("URL:", urlStr)` debug
  print inside `performSecureRequest` (uncommitted) — should be removed
  before merging.
- **`multi_source_route_summary.go` → `multi_source_summary.go`**: renamed;
  added `XRequestID` validation.

## Removed

- **`under_maintenance_func.go` deleted** (uncommitted), along with every
  `isUnderMaintenance(...)` guard previously in `Reverse`, `Search`,
  `SearchByRadius`, `SnapToRoad`, `DetailsByPlaceID`, and `Geocode`. This
  removes the ability to short-circuit those calls with a canned
  "service is under maintenance" response — confirm this is intentional
  before merging, since it's a behavioral removal, not just cleanup.

## Housekeeping / doc-comment removals (uncommitted)

- Removed the explanatory comments above `XRequestIDFromJSON` (`secure_request.go`),
  `ValidateAcceptedLanguage` (`validate_accepted_language.go`), and the
  now-commented-out `MultiSourceRouteSummaryRequestWithoutGeometry` struct.

## Suggested follow-ups before merging

1. Revert `baseURL` in `build_url_from_json.go` back to the production URL.
2. Remove the `fmt.Println("URL:", urlStr)` debug statement in `secure_request.go`.
3. Decide whether removing `under_maintenance_func.go` and all its call sites
   is intended for this PR, or should be split out.
4. Either delete or properly reinstate the commented-out
   `MultiSourceRouteSummaryRequestWithoutGeometry` struct in
   `multi_source_summary_without_geometry.go`.
5. Consider renaming `gecode.go` to `geocode.go` for consistency.
