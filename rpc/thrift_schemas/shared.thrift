namespace go shared
namespace java org.sharesound.android.rpc.shared

typedef string AuthToken

struct Session {
    1: required AuthToken auth_token;
}

exception AuthException {
    1: string message;
}