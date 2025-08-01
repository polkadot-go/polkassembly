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
	Wallet    string `json:"wallet"`
	Message   string `json:"message,omitempty"`
	Network   string `json:"network,omitempty"`
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
	Address   string `json:"address"`
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
	SearchTerm   string `json:"searchTerm,omitempty"`
}

type PostListingResponse struct {
	Posts      []Post `json:"posts"`
	Count      int    `json:"count"`
	TotalCount int    `json:"totalCount"`
}

type Post struct {
	PostID           int       `json:"post_id"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	Username         string    `json:"username"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PostType         string    `json:"post_type"`
	ProposalType     string    `json:"proposalType"`
	Status           string    `json:"status"`
	ProposerAddress  string    `json:"proposer"`
	CommentsCount    int       `json:"comments_count"`
	ReactionsCount   int       `json:"reactions_count"`
	ViewsCount       int       `json:"views_count"`
	Network          string    `json:"network"`
	TrackNumber      int       `json:"track_number"`
	Hash             string    `json:"hash"`
	Method           string    `json:"method"`
	MotionProposalId int       `json:"motion_proposal_id,omitempty"`
	BountyId         int       `json:"bounty_id,omitempty"`
	TipHash          string    `json:"tip_hash,omitempty"`
}

type PostOnchainData struct {
	Hash          string `json:"hash"`
	Status        string `json:"status"`
	AyesCount     int    `json:"ayesCount"`
	NaysCount     int    `json:"naysCount"`
	SupportAmount string `json:"supportAmount"`
	AgainstAmount string `json:"againstAmount"`
	Turnout       string `json:"turnout"`
	Electorate    string `json:"electorate"`
	Threshold     string `json:"threshold"`
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
	IsDeleted bool      `json:"is_deleted"`
}

type ActivityFeedItem struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Content   string      `json:"content"`
	PostID    int         `json:"post_id"`
	PostType  string      `json:"post_type"`
	Username  string      `json:"username"`
	CreatedAt time.Time   `json:"created_at"`
	Network   string      `json:"network"`
	Data      interface{} `json:"data,omitempty"`
}

type CreateOffchainPostRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	TopicID  int      `json:"topic_id"`
	Tags     []string `json:"tags,omitempty"`
	PostType string   `json:"post_type,omitempty"`
}

type UpdatePostRequest struct {
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

type SubscriptionStatus struct {
	Subscribed bool `json:"subscribed"`
}

type Bounty struct {
	BountyID       int    `json:"bounty_id"`
	Description    string `json:"description"`
	Proposer       string `json:"proposer"`
	Value          string `json:"value"`
	Fee            string `json:"fee"`
	Status         string `json:"status"`
	CuratorDeposit string `json:"curator_deposit,omitempty"`
	Bond           string `json:"bond,omitempty"`
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
	ID              string    `json:"id"`
	Voter           string    `json:"voter"`
	Balance         string    `json:"balance"`
	Vote            string    `json:"vote"`
	LockPeriod      int       `json:"lockPeriod"`
	Decision        string    `json:"decision"`
	CreatedAt       time.Time `json:"created_at"`
	DelegatedTo     string    `json:"delegatedTo,omitempty"`
	IsDelegated     bool      `json:"isDelegated"`
	ConvictionCount int       `json:"conviction_count"`
}

type VotingCurveData struct {
	BlockNumber int    `json:"blockNumber"`
	AyeAmount   string `json:"ayeAmount"`
	NayAmount   string `json:"nayAmount"`
	Support     string `json:"support"`
	Turnout     string `json:"turnout"`
}

type CreateVoteRequest struct {
	PostID     int    `json:"postId"`
	Vote       string `json:"vote"` // "aye" or "nay"
	Balance    string `json:"balance,omitempty"`
	LockPeriod int    `json:"lockPeriod,omitempty"`
}

// Action types
type AddCommentRequest struct {
	Content   string `json:"content"`
	PostID    int    `json:"postId"`
	PostType  string `json:"postType"`
	Sentiment int    `json:"sentiment,omitempty"`
	ParentID  string `json:"parentId,omitempty"`
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

type Report struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	ContentID  int       `json:"content_id"`
	Reason     string    `json:"reason"`
	Comments   string    `json:"comments"`
	ReportedBy string    `json:"reported_by"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
}

type CreateReportRequest struct {
	Type      string `json:"type"`
	ContentID int    `json:"content_id"`
	Reason    string `json:"reason"`
	Comments  string `json:"comments,omitempty"`
}

// User types
type User struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email,omitempty"`
	Web3Address    string    `json:"web3_address,omitempty"`
	EmailVerified  bool      `json:"email_verified"`
	Title          string    `json:"title,omitempty"`
	Bio            string    `json:"bio,omitempty"`
	Image          string    `json:"image,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	ProfileScore   int       `json:"profile_score"`
	FollowerCount  int       `json:"follower_count"`
	FollowingCount int       `json:"following_count"`
}

type UserActivity struct {
	Posts     []Post     `json:"posts"`
	Comments  []Comment  `json:"comments"`
	Reactions []Reaction `json:"reactions"`
	Votes     []Vote     `json:"votes"`
}

type UserListingParams struct {
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Sort  string `json:"sort,omitempty"`
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
	Author       string      `json:"author,omitempty"`
	Deposit      string      `json:"deposit,omitempty"`
}

type PreimageListingParams struct {
	Page   int    `json:"page,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Status string `json:"status,omitempty"`
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
	Image            string    `json:"image,omitempty"`
	Score            int       `json:"score"`
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

// Treasury types
type TreasuryProposal struct {
	ProposalID  int       `json:"proposal_id"`
	Proposer    string    `json:"proposer"`
	Value       string    `json:"value"`
	Beneficiary string    `json:"beneficiary"`
	Bond        string    `json:"bond"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTreasuryProposalRequest struct {
	Value       string `json:"value"`
	Beneficiary string `json:"beneficiary"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

// Tip types
type Tip struct {
	Hash      string    `json:"hash"`
	Who       string    `json:"who"`
	Finder    string    `json:"finder"`
	Reason    string    `json:"reason"`
	Status    string    `json:"status"`
	Tips      []TipInfo `json:"tips"`
	CreatedAt time.Time `json:"created_at"`
}

type TipInfo struct {
	Tipper string `json:"tipper"`
	Value  string `json:"value"`
}

type CreateTipRequest struct {
	Hash   string `json:"hash"`
	Reason string `json:"reason"`
	Who    string `json:"who"`
}

// Discussion types
type Discussion struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Author        string    `json:"author"`
	Tags          []string  `json:"tags"`
	ViewCount     int       `json:"view_count"`
	CommentCount  int       `json:"comment_count"`
	ReactionCount int       `json:"reaction_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LastCommentAt time.Time `json:"last_comment_at,omitempty"`
}

type CreateDiscussionRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags,omitempty"`
}

// Poll types
type Poll struct {
	ID         int          `json:"id"`
	Question   string       `json:"question"`
	Options    []PollOption `json:"options"`
	EndAt      time.Time    `json:"end_at"`
	CreatedBy  string       `json:"created_by"`
	CreatedAt  time.Time    `json:"created_at"`
	VoterCount int          `json:"voter_count"`
	Status     string       `json:"status"`
}

type PollOption struct {
	ID         int     `json:"id"`
	Text       string  `json:"text"`
	VoteCount  int     `json:"vote_count"`
	Percentage float64 `json:"percentage"`
}

type CreatePollRequest struct {
	Question string    `json:"question"`
	Options  []string  `json:"options"`
	EndAt    time.Time `json:"end_at"`
	PostID   int       `json:"post_id,omitempty"`
}

type PollVoteRequest struct {
	OptionID int `json:"option_id"`
}

// Notification types
type Notification struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	URL       string    `json:"url,omitempty"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationPreferences struct {
	NewProposal          bool `json:"new_proposal"`
	ProposalStatusChange bool `json:"proposal_status_change"`
	NewComment           bool `json:"new_comment"`
	NewReaction          bool `json:"new_reaction"`
	NewFollower          bool `json:"new_follower"`
	NewMention           bool `json:"new_mention"`
}

// Analytics types
type ProposalAnalytics struct {
	PostID         int            `json:"post_id"`
	ViewCount      int            `json:"view_count"`
	UniqueViewers  int            `json:"unique_viewers"`
	CommentCount   int            `json:"comment_count"`
	ReactionCount  int            `json:"reaction_count"`
	VoteCount      int            `json:"vote_count"`
	ShareCount     int            `json:"share_count"`
	DailyStats     []DailyStat    `json:"daily_stats"`
	VoterBreakdown map[string]int `json:"voter_breakdown"`
	TopCommenters  []UserStat     `json:"top_commenters"`
}

type DailyStat struct {
	Date      time.Time `json:"date"`
	Views     int       `json:"views"`
	Comments  int       `json:"comments"`
	Votes     int       `json:"votes"`
	Reactions int       `json:"reactions"`
}

type UserStat struct {
	Username string `json:"username"`
	Count    int    `json:"count"`
}

// Network Stats
type NetworkStats struct {
	ActiveProposals       int    `json:"active_proposals"`
	TotalProposals        int    `json:"total_proposals"`
	TotalVotes            int    `json:"total_votes"`
	TotalUsers            int    `json:"total_users"`
	TotalDelegations      int    `json:"total_delegations"`
	TotalDelegatedBalance string `json:"total_delegated_balance"`
	WeeklyActiveUsers     int    `json:"weekly_active_users"`
	MonthlyActiveUsers    int    `json:"monthly_active_users"`
}

// Search types
type SearchParams struct {
	Query    string    `json:"query"`
	Type     string    `json:"type,omitempty"` // "all", "posts", "comments", "users"
	Network  string    `json:"network,omitempty"`
	Author   string    `json:"author,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
	DateFrom time.Time `json:"date_from,omitempty"`
	DateTo   time.Time `json:"date_to,omitempty"`
	Status   string    `json:"status,omitempty"`
	TrackNo  int       `json:"track_no,omitempty"`
	Page     int       `json:"page,omitempty"`
	Limit    int       `json:"limit,omitempty"`
}

type SearchResponse struct {
	Posts      []Post    `json:"posts,omitempty"`
	Comments   []Comment `json:"comments,omitempty"`
	Users      []User    `json:"users,omitempty"`
	TotalCount int       `json:"total_count"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}

// Timeline types
type TimelineEntry struct {
	ID          int         `json:"id"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	Author      string      `json:"author"`
	CreatedAt   time.Time   `json:"created_at"`
	Data        interface{} `json:"data"`
	Network     string      `json:"network"`
	IsImportant bool        `json:"is_important"`
}

type TimelineParams struct {
	Network  string    `json:"network,omitempty"`
	DateFrom time.Time `json:"date_from,omitempty"`
	DateTo   time.Time `json:"date_to,omitempty"`
	Types    []string  `json:"types,omitempty"`
	Page     int       `json:"page,omitempty"`
	Limit    int       `json:"limit,omitempty"`
}
