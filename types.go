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
	Network   string `json:"network"`
}

type Web3AuthResponse struct {
	Token   string `json:"token"`
	User    User   `json:"user"`
	Message string `json:"message,omitempty"`
}

type Web2LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Web2LoginResponse struct {
	Token   string `json:"token"`
	User    User   `json:"user"`
	Message string `json:"message,omitempty"`
}

type Web2SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Web2SignupResponse struct {
	Token   string `json:"token"`
	User    User   `json:"user"`
	Message string `json:"message,omitempty"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type QRSessionResponse struct {
	SessionID string `json:"sessionId"`
	QRCode    string `json:"qrCode"`
}

type ClaimQRSessionRequest struct {
	SessionID string `json:"sessionId"`
	Signature string `json:"signature"`
}

type EditUserDetailsRequest struct {
	Username    string   `json:"username,omitempty"`
	Email       string   `json:"email,omitempty"`
	Bio         string   `json:"bio,omitempty"`
	Image       string   `json:"image,omitempty"`
	Title       string   `json:"title,omitempty"`
	SocialLinks []string `json:"social_links,omitempty"`
}

// Post types
type PostListingParams struct {
	Page         int    `json:"page,omitempty"`
	ListingLimit int    `json:"listingLimit,omitempty"`
	TrackNo      int    `json:"trackNo,omitempty"`
	TrackStatus  string `json:"trackStatus,omitempty"`
	ProposalType string `json:"proposalType,omitempty"`
	SortBy       string `json:"sortBy,omitempty"`
}

type PostListingResponse struct {
	Posts []Post `json:"posts"`
	Count int    `json:"count"`
}

type Post struct {
	PostID          int       `json:"post_id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Username        string    `json:"username"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PostType        string    `json:"post_type"`
	ProposalType    string    `json:"proposalType"`
	Status          string    `json:"status"`
	ProposerAddress string    `json:"proposer"`
	CommentsCount   int       `json:"comments_count"`
	ReactionsCount  int       `json:"reactions_count"`
	Network         string    `json:"network"`
	TrackNumber     int       `json:"track_number"`
}

type PostOnchainData struct {
	Hash          string `json:"hash"`
	Status        string `json:"status"`
	AyesCount     int    `json:"ayesCount"`
	NaysCount     int    `json:"naysCount"`
	SupportAmount string `json:"supportAmount"`
	AgainstAmount string `json:"againstAmount"`
}

type ContentSummary struct {
	Summary  string   `json:"summary"`
	Keywords []string `json:"keywords"`
}

type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Replies   []Comment `json:"replies,omitempty"`
	Sentiment int       `json:"sentiment"`
}

type ActivityFeedItem struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	PostID    int       `json:"post_id"`
	PostType  string    `json:"post_type"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Network   string    `json:"network"`
}

type CreateOffchainPostRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	TopicID int      `json:"topic_id"`
	Tags    []string `json:"tags,omitempty"`
}

type UpdatePostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type SubscriptionStatus struct {
	Subscribed bool `json:"subscribed"`
}

type Bounty struct {
	BountyID    int    `json:"bounty_id"`
	Description string `json:"description"`
	Proposer    string `json:"proposer"`
	Value       string `json:"value"`
	Fee         string `json:"fee"`
	Status      string `json:"status"`
}

// Vote types
type VoteListingParams struct {
	PostID   int    `json:"postId,omitempty"`
	Page     int    `json:"page,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	VoteType string `json:"voteType,omitempty"`
}

type VoteListingResponse struct {
	Votes []Vote `json:"votes"`
	Count int    `json:"count"`
}

type Vote struct {
	ID          string    `json:"id"`
	Voter       string    `json:"voter"`
	Balance     string    `json:"balance"`
	Vote        string    `json:"vote"`
	LockPeriod  int       `json:"lockPeriod"`
	Decision    string    `json:"decision"`
	CreatedAt   time.Time `json:"created_at"`
	DelegatedTo string    `json:"delegatedTo,omitempty"`
	IsDelegated bool      `json:"isDelegated"`
}

type VotingCurveData struct {
	BlockNumber int    `json:"blockNumber"`
	AyeAmount   string `json:"ayeAmount"`
	NayAmount   string `json:"nayAmount"`
	Support     string `json:"support"`
}

// Action types
type AddCommentRequest struct {
	Content   string `json:"content"`
	PostID    int    `json:"postId"`
	PostType  string `json:"postType"`
	Sentiment int    `json:"sentiment,omitempty"`
}

type AddReactionRequest struct {
	PostID    int    `json:"postId"`
	PostType  string `json:"postType"`
	Reaction  string `json:"reaction"`
	CommentID string `json:"commentId,omitempty"`
}

type UpdateCommentRequest struct {
	Content   string `json:"content"`
	Sentiment int    `json:"sentiment,omitempty"`
}

type Reaction struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Reaction  string    `json:"reaction"`
	CreatedAt time.Time `json:"created_at"`
}

// User types
type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email,omitempty"`
	Web3Address   string    `json:"web3_address,omitempty"`
	EmailVerified bool      `json:"email_verified"`
	Title         string    `json:"title,omitempty"`
	Bio           string    `json:"bio,omitempty"`
	Image         string    `json:"image,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

type UserActivity struct {
	Posts     []Post     `json:"posts"`
	Comments  []Comment  `json:"comments"`
	Reactions []Reaction `json:"reactions"`
}

type UserListingParams struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type UserListingResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

// Preimage types
type Preimage struct {
	Hash         string      `json:"hash"`
	Length       int         `json:"length"`
	Method       string      `json:"method"`
	Section      string      `json:"section"`
	ProposedCall interface{} `json:"proposedCall"`
	Status       string      `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
}

type PreimageListingParams struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type PreimageListingResponse struct {
	Preimages []Preimage `json:"preimages"`
	Count     int        `json:"count"`
}

// Vote Cart types
type CartItem struct {
	ID           string    `json:"id"`
	PostID       int       `json:"postId"`
	ProposalType string    `json:"proposalType"`
	AyeBalance   string    `json:"ayeBalance"`
	NayBalance   string    `json:"nayBalance"`
	Abstain      string    `json:"abstain"`
	LockPeriod   int       `json:"lockPeriod"`
	CreatedAt    time.Time `json:"created_at"`
}

type AddCartItemRequest struct {
	PostID       int    `json:"postId"`
	ProposalType string `json:"proposalType"`
	AyeBalance   string `json:"ayeBalance,omitempty"`
	NayBalance   string `json:"nayBalance,omitempty"`
	Abstain      string `json:"abstain,omitempty"`
	LockPeriod   int    `json:"lockPeriod,omitempty"`
}

type UpdateCartItemRequest struct {
	AyeBalance string `json:"ayeBalance,omitempty"`
	NayBalance string `json:"nayBalance,omitempty"`
	Abstain    string `json:"abstain,omitempty"`
	LockPeriod int    `json:"lockPeriod,omitempty"`
}

// Delegation types
type DelegationStats struct {
	TotalDelegations  int    `json:"totalDelegations"`
	TotalDelegates    int    `json:"totalDelegates"`
	TotalBalance      string `json:"totalBalance"`
	WeeklyDelegations int    `json:"weeklyDelegations"`
}

type Delegate struct {
	Address          string    `json:"address"`
	Name             string    `json:"name"`
	Bio              string    `json:"bio"`
	DelegationsCount int       `json:"delegations_count"`
	ActiveProposals  int       `json:"active_proposals"`
	VotedProposals   int       `json:"voted_proposals"`
	CreatedAt        time.Time `json:"created_at"`
}

type CreatePADelegateRequest struct {
	Address string   `json:"address"`
	Name    string   `json:"name"`
	Bio     string   `json:"bio"`
	Tags    []string `json:"tags,omitempty"`
}

type UpdatePADelegateRequest struct {
	Name string   `json:"name,omitempty"`
	Bio  string   `json:"bio,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

type TrackStats struct {
	TrackID          int    `json:"trackId"`
	TrackName        string `json:"trackName"`
	DelegatedAmount  string `json:"delegatedAmount"`
	DelegationsCount int    `json:"delegationsCount"`
}

type TrackLevelData struct {
	TrackID    int `json:"trackId"`
	Level      int `json:"level"`
	Multiplier int `json:"multiplier"`
}
