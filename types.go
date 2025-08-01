package polkassembly

import "time"

// Common types
type APIError struct {
	ErrorMessage string `json:"error"`
	Message      string `json:"message"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.ErrorMessage
}

// Auth types
type Web3AuthRequest struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
	Message   string `json:"message"`
}

type Web3AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Web2LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Web2LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Web2SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Web2SignupResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type QRSessionResponse struct {
	SessionID string `json:"session_id"`
	QRCode    string `json:"qr_code"`
}

type ClaimQRSessionRequest struct {
	SessionID string `json:"session_id"`
	Token     string `json:"token"`
}

type EditUserDetailsRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

// Post types
type PostListingParams struct {
	Page            int    `json:"page,omitempty"`
	Limit           int    `json:"limit,omitempty"`
	Type            string `json:"type,omitempty"`
	Status          string `json:"status,omitempty"`
	ProposerAddress string `json:"proposer_address,omitempty"`
}

type PostListingResponse struct {
	Posts      []Post `json:"posts"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

type Post struct {
	ID              int        `json:"id"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	Author          User       `json:"author"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Type            string     `json:"type"`
	Status          string     `json:"status"`
	OnchainID       string     `json:"onchain_id"`
	ProposerAddress string     `json:"proposer_address"`
	Comments        []Comment  `json:"comments,omitempty"`
	Reactions       []Reaction `json:"reactions,omitempty"`
}

type PostOnchainData struct {
	BlockNumber    int    `json:"block_number"`
	Hash           string `json:"hash"`
	Status         string `json:"status"`
	AyeVotes       string `json:"aye_votes"`
	NayVotes       string `json:"nay_votes"`
	TurnoutPercent string `json:"turnout_percent"`
}

type ContentSummary struct {
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

type Comment struct {
	ID        int        `json:"id"`
	Content   string     `json:"content"`
	Author    User       `json:"author"`
	PostID    int        `json:"post_id"`
	ParentID  *int       `json:"parent_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Reactions []Reaction `json:"reactions,omitempty"`
}

type ActivityFeedItem struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Action    string    `json:"action"`
	Actor     User      `json:"actor"`
	Target    string    `json:"target"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateOffchainPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type UpdatePostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type SubscriptionStatus struct {
	IsSubscribed bool `json:"is_subscribed"`
}

type Bounty struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       string `json:"value"`
	Status      string `json:"status"`
}

// Vote types
type VoteListingParams struct {
	PostID   int    `json:"post_id,omitempty"`
	Page     int    `json:"page,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Decision string `json:"decision,omitempty"`
}

type VoteListingResponse struct {
	Votes      []Vote `json:"votes"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

type Vote struct {
	ID        int       `json:"id"`
	Address   string    `json:"address"`
	Balance   string    `json:"balance"`
	Decision  string    `json:"decision"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type VotingCurveData struct {
	Timestamp time.Time `json:"timestamp"`
	AyeVotes  string    `json:"aye_votes"`
	NayVotes  string    `json:"nay_votes"`
	Turnout   string    `json:"turnout"`
}

// Action types
type AddCommentRequest struct {
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	ParentID *int   `json:"parent_id,omitempty"`
}

type AddReactionRequest struct {
	PostID    int    `json:"post_id,omitempty"`
	CommentID int    `json:"comment_id,omitempty"`
	Reaction  string `json:"reaction"`
}

type UpdateCommentRequest struct {
	Content string `json:"content"`
}

type Reaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
	Reaction  string    `json:"reaction"`
	CreatedAt time.Time `json:"created_at"`
}

// User types
type User struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email,omitempty"`
	Address        string    `json:"address,omitempty"`
	Bio            string    `json:"bio,omitempty"`
	Avatar         string    `json:"avatar,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
}

type UserActivity struct {
	Activities []ActivityFeedItem `json:"activities"`
	TotalCount int                `json:"total_count"`
	Page       int                `json:"page"`
	Limit      int                `json:"limit"`
}

type UserListingParams struct {
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Sort  string `json:"sort,omitempty"`
}

type UserListingResponse struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

// Preimage types
type Preimage struct {
	Hash       string    `json:"hash"`
	Length     int       `json:"length"`
	Data       string    `json:"data"`
	ProposalID int       `json:"proposal_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type PreimageListingParams struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type PreimageListingResponse struct {
	Preimages  []Preimage `json:"preimages"`
	TotalCount int        `json:"total_count"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
}

// Vote Cart types
type CartItem struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Decision  string    `json:"decision"`
	Balance   string    `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type AddCartItemRequest struct {
	PostID   int    `json:"post_id"`
	Decision string `json:"decision"`
	Balance  string `json:"balance"`
}

type UpdateCartItemRequest struct {
	Decision string `json:"decision,omitempty"`
	Balance  string `json:"balance,omitempty"`
}

// Delegation types
type DelegationStats struct {
	TotalDelegates    int    `json:"total_delegates"`
	TotalDelegations  int    `json:"total_delegations"`
	TotalAmount       string `json:"total_amount"`
	ActiveDelegations int    `json:"active_delegations"`
}

type Delegate struct {
	ID               int       `json:"id"`
	Address          string    `json:"address"`
	Name             string    `json:"name"`
	Bio              string    `json:"bio"`
	DelegationsCount int       `json:"delegations_count"`
	TotalBalance     string    `json:"total_balance"`
	CreatedAt        time.Time `json:"created_at"`
}

type CreatePADelegateRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type UpdatePADelegateRequest struct {
	Name string `json:"name,omitempty"`
	Bio  string `json:"bio,omitempty"`
}

type TrackStats struct {
	TrackID           int    `json:"track_id"`
	TrackName         string `json:"track_name"`
	Delegations       int    `json:"delegations"`
	TotalBalance      string `json:"total_balance"`
	VoteParticipation string `json:"vote_participation"`
}

type TrackLevelData struct {
	TrackID   int    `json:"track_id"`
	TrackName string `json:"track_name"`
	Level     int    `json:"level"`
	Points    int    `json:"points"`
	NextLevel int    `json:"next_level"`
}
