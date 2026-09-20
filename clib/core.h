#ifndef MAPNESTS_CORE_H
#define MAPNESTS_CORE_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    int   success;
    int   status_code;
    char *response;
    char *error_message;
} SecureResult;

char        *build_url_from_json(const char *label, const char *json_str);
SecureResult perform_secure_request(const char *label, const char *api_key,
                                    const char *origin, int timeout_ms,
                                    const char *json_request,
                                    const char *x_request_id);
void         free_secure_result(SecureResult *r);
void         free_string(char *s);

#ifdef __cplusplus
}
#endif

#endif /* MAPNESTS_CORE_H */
