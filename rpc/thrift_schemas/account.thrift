include "shared.thrift"

namespace go account
namespace java org.sharesound.android.rpc.account

struct ProfileResult {
    1: required shared.Session session;

    2: required string email;
    3: required string username;
}

service AccountService {

    shared.Session registerAccount(1:string email,
                                    2:string username,
                                    3:string password),

    shared.Session login(1:string email, 2:string password) throws (1:shared.AuthException auth_e),

    void logout(1:shared.Session ctx) throws (1:shared.AuthException auth_e);

    ProfileResult getProfile(1:shared.Session ctx) throws (1:shared.AuthException auth_e);
}