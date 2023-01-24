namespace go api

struct User {
    1: i64 id // User id
    2: string name // Username
    3: i64 follow_count // Total number of followings
    4: i64 follower_count // Total number of followers
    5: bool is_follow // true-followed, false-not followed
}

struct Video {
    1: i64 id // Video unique identifier
    2: User author // Video author information
    3: string play_url // Video play URL
    4: string cover_url // Video cover address
    5: i64 favorite_count // Total number of likes for the video
    6: i64 comment_count // Total number of comments on the video
    7: bool is_favorite // true-liked, false-not liked
    8: string title // Video title
}

struct Comment {
    1: i64 id // Video comment id
    2: User user // Comment user information
    3: string content // Comment's content
    4: string create_date // Comment release date, format mm-dd
}

struct douyin_user_register_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // Username, up to 32 characters
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // Password, up to 32 characters
}

struct douyin_user_register_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: i64 user_id // User id
    4: string token // User authentication token
}

struct douyin_user_login_request {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // Username, up to 32 characters
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // Password, up to 32 characters
}

struct douyin_user_login_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: i64 user_id // User id
    4: string token // User authentication token
}

struct douyin_user_request {
    1: i64 user_id(api.query="user_id", api.vd="len($)>0 && len($)<20") // User id
    2: string token(api.query="token") // User authentication token
}

struct douyin_user_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: User user // User Information
}

struct douyin_feed_request {
    1: i64 latest_time(api.query="latest_time") // Optional parameter, limit the latest submission timestamp of the returned video, accurate to seconds, and leave it blank to indicate the current time
    2: string token(api.query="token") // Optional parameter, login user settings
}

struct douyin_feed_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: Video video_list // Video list
    4: i64 next_time // In the video returned this time, publish the earliest time as the latest_time in the next request
}

struct douyin_publish_action_request {
    1: string token(api.form="token") // User authentication token
    2: byte data(api.form="data") // Video data
    3: string title(api.form="title", api.vd="len($)>0 && len($)<33") // Video title
}

struct douyin_publish_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct douyin_publish_list_request {
    1: i64 user_id(api.query="user_id", api.vd="len($)>0 && len($)<20") // User id
    2: string token(api.query="token") // User authentication token
}

struct douyin_publish_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<Video> video_list // List of videos posted by the user
}

struct douyin_favorite_action_request {
    1: string token(api.query="token") // User authentication token
    2: i64 video_id(api.query="video_id", api.vd="len($)>0 && len($)<20") // Video Id
    3: i32 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-like, 2-unlike
}

struct douyin_favorite_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct douyin_favorite_list_request {
    1: i64 user_id(api.query="user_id", api.vd="len($)>0 && len($)<20") // User id
    2: string token(api.query="token") // User authentication token
}

struct douyin_favorite_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<Video> video_list // List of videos posted by the user
}

struct douyin_comment_action_request {
    1: string token(api.query="token") // User authentication token
    2: i64 video_id(api.query="video_id", api.vd="len($)>0 && len($)<20") // Video Id
    3: i32 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-like, 2-unlike
    4: string comment_text(api.query="comment_text") // The content of the comment filled by the user, used when action_type=1
    5: i64 comment_id(api.query="comment_id", api.vd="len($)>0 && len($)<20") // The comment id to be deleted is used when action_type=2
}

struct douyin_comment_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: Comment comment // The comment successfully returns the comment content, no need to re-pull the entire list
}

struct douyin_comment_list_request {
    1: string token(api.query="token") // User authentication token
    2: i64 video_id(api.query="video_id", api.vd="len($)>0 && len($)<20") // Video Id
}

struct douyin_comment_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<Comment> comment_list // List of comments
}

struct douyin_relation_action_request {
    1: string token(api.query="token") // User authentication token
    2: i64 to_user_id(api.query="to_user_id", api.vd="len($)>0 && len($)<20") // The other party's user id
    3: i32 action_type(api.query="action_type", api.vd="$==1 || $==2") // 1-Follow, 2-Unfollow
}

struct douyin_relation_action_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
}

struct douyin_relation_follow_list_request {
    1: i64 user_id(api.query="user_id", api.vd="len($)>0 && len($)<20") // User id
    2: string token(api.query="token") // User authentication token
}

struct douyin_relation_follow_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<User> user_list // List of user information
}

struct douyin_relation_follower_list_request {
    1: i64 user_id(api.query="user_id", api.vd="len($)>0 && len($)<20") // User id
    2: string token(api.query="token") // User authentication token
}

struct douyin_relation_follower_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<User> user_list // List of user information
}

struct douyin_relation_friend_list_request {
    1: i64 user_id(api.query="user_id", api.vd="len($)>0 && len($)<20") // User id
    2: string token(api.query="token") // User authentication token
}

struct douyin_relation_friend_list_response {
    1: i32 status_code // Status code, 0-success, other values-failure
    2: string status_msg // Return status description
    3: list<User> user_list // List of user information
}

service ApiService {
    douyin_user_register_response Register(1: douyin_user_register_request req)(api.post="/douyin/user/register");
    douyin_user_login_response Login(1: douyin_user_login_request req)(api.post="/douyin/user/login");
    douyin_user_request GetUserInfo(1: douyin_user_request req)(api.get="/douyin/user");

    douyin_feed_response Feed (1: douyin_feed_request req)(api.get="/douyin/feed");
    douyin_publish_action_response PublishVideo (1: douyin_publish_action_request req)(api.post="/douyin/publish/action");
    douyin_publish_list_response VideoList (1: douyin_publish_list_request req)(api.get="/douyin/publish/list");

    douyin_favorite_action_response Favorite(1: douyin_favorite_action_request req)(api.post="/douyin/favorite/action");
    douyin_favorite_list_response FavoriteList(1: douyin_favorite_list_request req)(api.get="/douyin/favorite/list");
    douyin_comment_action_response Comment(1: douyin_comment_action_request req)(api.post="/douyin/comment/action");
    douyin_comment_list_response CommentList(1: douyin_comment_list_request req)(api.get="/douyin/comment/list");

    douyin_relation_action_response Action(1: douyin_relation_action_request req)(api.post="/douyin/relation/action");
    douyin_relation_follow_list_response FollowingList(1: douyin_relation_follow_list_request req)(api.get="/douyin/relation/follow/list");
    douyin_relation_follower_list_response FollowerList(1: douyin_relation_follower_list_request req)(api.get="/douyin/relation/follower/list");
    douyin_relation_friend_list_response FriendList(1: douyin_relation_friend_list_request req)(api.get="/douyin/relation/friend/list");
}