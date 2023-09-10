package docs

///// account-rules ==>> table
// Transaction Limits:

// Daily/Weekly/Monthly Transaction Limits: Restrict the number of transactions a user can make within specific time frames.
// Minimum and Maximum Amounts: Define the smallest and largest amounts for a single transaction.
// Velocity Checks:

// Multiple Failed Transactions: If a user has multiple failed transactions within a short period, it might be an indicator of fraud or system abuse.
// Rapid, High-Frequency Transactions: If a user is making many transactions in quick succession, it could be suspicious.
// Suspicious Activity:

// Unusual Hours: Transactions made during non-typical hours might be considered higher risk.
// Rapid Changes in Transaction Amounts: If a user usually transacts small amounts but suddenly transacts a large amount, it might warrant a review.
// Location-Based Checks:

// Geographical Restrictions: Some platforms restrict transactions based on the user's geographical location.
// Unusual Locations: If a user typically makes transactions from one location but suddenly starts transacting from a very different location, it might be considered suspicious.
// User Behavior:

// New Device or IP: A transaction from a new device or IP might be treated with caution, especially if the amount is large.
// Change of Personal Information: If personal details (like email, phone number) change followed by a transaction, it might be considered suspicious.
// Account Age:

// New Accounts: New accounts might have stricter rules initially. For instance, they might have lower transaction limits, and these limits can be raised as the account "ages" and establishes trust.
// Account Balance Restrictions:

// Insufficient Funds: Obviously, transactions should be checked against the available balance.
// Minimum Balance Requirement: Some accounts might require a minimum balance to be maintained.
// Authentication and Authorization:

// Two-Factor Authentication (2FA): For high-value transactions, requiring 2FA can add an additional layer of security.
// Blacklists and Whitelists:

// Blacklist: A list of accounts, IP addresses, or devices that are forbidden from making transactions due to previous suspicious activities.
// Whitelist: A list of trusted entities which are always allowed to transact.
// Tiered Verification:

// Some platforms adopt a tiered verification approach. Users who provide minimal verification might have lower limits, and as they provide more documentation or verification (like linking a bank account, verifying a phone number), their limits are raised.
// Cooling Periods:
// After certain actions, like changing a password or adding a new transaction method, you might enforce a waiting period before transactions can be made.

// User-specific vs. Global Rules: The above design applies rules on a per-user basis. If you want global rules, you might move away from linking to a specific user_id.

// More Complex Rules: Some rules are not easily captured in a table column. For example, checking for transactions at unusual hours would likely be an application-level check rather than a database rule.

// Blacklists and Whitelists: These might be better suited for their own tables. For instance, a blacklisted_accounts table or a whitelisted_accounts table.

// Location-based Checks: These also might be handled better at the application level, using geolocation APIs and user profile data.

// Normalization: For cleaner design, some columns, like verification_level, might be better served in normalized, separate tables.

// Dynamic Rule Addition: If you need to frequently add new types of rules, you might consider a more flexible schema or even a rule engine.

// When designing tables to hold rules, it's vital to anticipate changes. The regulatory environment, business needs, and fraud patterns all evolve, so your rule system must be adaptable.

// If every transaction will always be in the currency of the associated financial account, then you don't need to redundantly store the currency in each transaction table (card_transaction and account_transaction). You can retrieve the currency by joining the transaction table with the financial_account table and then with the currency table based on the account_id.

// However, there are reasons why you might still want to include the currency_id in your transaction tables:

// Performance: Joining tables can have a performance impact, especially as the database grows. Storing the currency_id directly in the transaction tables can save you from having to perform joins every time you want to retrieve the currency of a transaction, especially if this is a frequent operation.

// Historical Accuracy: If, for some reason, an account's currency can change (though this is uncommon), having the currency_id in the transaction table ensures that you have an accurate historical record of the currency of each transaction.

// Transactions in Different Currencies: If a card or account transaction can occur in a different currency than the default currency of the financial account, then you will definitely need to store the currency_id in the transaction tables.

// If none of these reasons apply to your use case, and every transaction is guaranteed to be in the currency of the associated financial account, then you can avoid storing the currency_id in the transaction tables and retrieve it from the financial_account table when needed.

// Transfer initiates a transfer between accounts
// func (s *Service) Transfer(ctx context.Context, req request.TransferRequest) (res response.TransferResponse, err error) {
// s.logger.Info("Starting the process to transfer funds",
// 	zap.Int("SenderAccountID", req.SenderAccountID),
// 	zap.Int("ReceiverAccountID", req.ReceiverAccountID),
// 	zap.Float64("Amount", req.Amount))

// Potential Improvements
// Idempotency: In distributed systems, you could end up with duplicate requests. Making your transfer operation idempotent would prevent unintended duplicate transfers.

// Event Sourcing: Consider using an event-sourced architecture for even more robustness. Each transaction would be an "event" that alters the system's state.

// Logs and Monitoring: You already have logging, but ensure you're also monitoring these logs for anomalies.

// Graceful Error Handling: Provide specific error codes and messages for different types of failure - insufficient funds, account not found, etc.

// Concurrency: Ensure that your database operations are concurrency-safe. You've already used transactions, but you could also use database-level locking mechanisms to prevent other operations from intervening during a transfer.

// Validation: Besides request validation, you may also consider business rule validations like maximum/minimum transfer limits, account status checks, etc.

// Rate Limiting: To protect against abuse, consider implementing rate limiting on transfer requests.

// Notifications: Consider notifying the customer or internal teams during critical failures so that corrective measures can be immediately taken.

// Testing: Last but not the least, have a thorough set of unit tests and integration tests to ensure the business logic is working as intended.







// ***************************************************************************************************************
// populating the account-transaction-status:
// Decide the transaction status

// For transactions that exceed a certain value, you might want to set the status to PendingApproval until a human manually approves the transaction.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if req.Amount > 10000 {  // 10,000 is the threshold for manual approval
// 	transactionStatus = enum.PendingApproval // Awaiting manual approval for high-value transactions
// } else {
// 	transactionStatus = enum.Completed // Lower-value transactions are automatically completed
// }

// 2. Third-Party Payment Gateway Confirmation
// If you are using a third-party payment gateway, you might want to set the transaction status to Pending until you get a confirmation from the payment gateway.
// // Communicate with the third-party payment gateway
// isConfirmed, err := s.paymentGatewayService.ConfirmPayment(ctx, req)
// if err != nil {
// 	s.logger.Error("Failed to confirm payment with third-party gateway", zap.Error(err))
// 	return res, err
// }
// // Decide the transaction status based on third-party confirmation
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isConfirmed {
// 	transactionStatus = enum.Completed // Transaction confirmed by the payment gateway
// } else {
// 	transactionStatus = enum.Pending // Pending until the payment gateway confirms
// }

// 3. Transaction During Non-Banking Hours
// Some banking systems do not complete transactions outside of banking hours, instead putting them in a "Pending" state.
// // Get the current time in UTC
// currentUTC := time.Now().UTC()
// hour := currentUTC.Hour()
// // Decide the transaction status based on the time
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if hour >= 9 && hour <= 17 { // Transactions are only auto-completed between 9 AM and 5 PM UTC
// 	transactionStatus = enum.Completed
// } else {
// 	transactionStatus = enum.Pending // Pending if outside of banking hours
// }

// 4. Fraud Detection
// If a fraud detection system is in place, you may want to mark the transaction as PendingReview if the system flags it for any reason.
// isSuspicious, err := s.fraudDetectionService.CheckTransaction(ctx, req)
// if err != nil {
// 	s.logger.Error("Failed to run fraud detection", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isSuspicious {
// 	transactionStatus = enum.PendingReview // Pending manual review due to suspicious activity
// } else {
// 	transactionStatus = enum.Completed // Transaction seems legitimate
// }

// 5. Special Holiday Rules
// Some banks have special processing rules during holidays. Transactions may not be completed instantly on holidays.
// isHoliday, err := s.calendarService.IsBankHoliday(currentUTC)
// if err != nil {
// 	s.logger.Error("Failed to check bank holiday", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isHoliday {
// 	transactionStatus = enum.Pending // Pending if it's a bank holiday
// } else {
// 	transactionStatus = enum.Completed
// }

// 6. Cross-Border Transactions
// For transactions that are happening across borders, you may want to set them to a different status until all international checks are completed.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isInternationalTransaction(req.SenderAccountID, req.ReceiverAccountID) { // You define what this function does
// 	transactionStatus = enum.PendingInternationalChecks
// } else {
// 	transactionStatus = enum.Completed
// }

// 7. User-Defined Controls
// Allow the user to set their own rules for auto-approving transactions, which you then evaluate.
// autoApprovalSetting, err := s.userSettingsService.GetAutoApprovalSetting(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to fetch user settings", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if autoApprovalSetting {
// 	transactionStatus = enum.Completed // User has enabled auto-approval
// } else {
// 	transactionStatus = enum.PendingUserApproval // Pending until user approves
// }

// 8. VIP Customers
// You may have VIP or premium customers who have additional features like instant transfers, even in scenarios where other users might be limited.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// isVIP, err := s.customerService.IsVIP(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to check if sender is a VIP customer", zap.Error(err))
// 	return res, err
// }

// if isVIP {
// 	transactionStatus = enum.Completed // Instant processing for VIP customers
// } else {
// 	transactionStatus = enum.Pending
// }

// 9. Scheduled Transfers
// For scheduled transfers, you may want to set the status as Scheduled until the transaction is actually processed.
// if req.IsScheduledTransfer {
// 	transactionStatus = enum.Scheduled
// } else {
// 	transactionStatus = enum.Completed
// }

// 10. Large Transactions
// For transactions that exceed a certain amount, you may need additional verification or manual approval.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if req.Amount > LARGE_AMOUNT_THRESHOLD { // Replace LARGE_AMOUNT_THRESHOLD with an actual number
// 	transactionStatus = enum.PendingManualApproval
// } else {
// 	transactionStatus = enum.Completed
// }

// 11. Multiple Failed Attempts
// If a user has had multiple failed attempts in a short period, flag their next transactions for review.
// failedAttempts, err := s.transactionService.RecentFailedAttempts(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to fetch recent failed attempts", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if failedAttempts > MAX_FAILED_ATTEMPTS { // Replace MAX_FAILED_ATTEMPTS with an actual number
// 	transactionStatus = enum.PendingReview
// } else {
// 	transactionStatus = enum.Completed
// }

// 12. Weekend Processing
// Some financial institutions dont process transactions over the weekend. You could use a service to determine if the current day is a weekend.
// isWeekend, err := s.calendarService.IsWeekend(currentUTC)
// if err != nil {
// 	s.logger.Error("Failed to check if today is a weekend", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isWeekend {
// 	transactionStatus = enum.Pending // Pending if it's a weekend
// } else {
// 	transactionStatus = enum.Completed
// }

// 13. Holiday Processing
// In some cases, transactions might not be processed on holidays. You could have a holiday calendar service to check this.
// isHoliday, err := s.calendarService.IsHoliday(currentUTC)
// if err != nil {
// 	s.logger.Error("Failed to check if today is a holiday", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isHoliday {
// 	transactionStatus = enum.Pending
// } else {
// 	transactionStatus = enum.Completed
// }

// 14. Fraud Detection
// If the transaction matches certain fraud detection criteria, it may be flagged for manual review.
// isSuspicious, err := s.fraudDetectionService.IsTransactionSuspicious(req)
// if err != nil {
// 	s.logger.Error("Failed to check for fraud detection", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isSuspicious {
// 	transactionStatus = enum.PendingReview
// } else {
// 	transactionStatus = enum.Completed
// }

// 15. Special Promotions or Campaigns
// During special promotions or campaigns, transactions might be auto-completed or given preferential treatment.
// isSpecialPromotion, err := s.promotionService.IsSpecialPromotion(currentUTC)
// if err != nil {
// 	s.logger.Error("Failed to check for special promotion", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isSpecialPromotion {
// 	transactionStatus = enum.Completed
// } else {
// 	transactionStatus = enum.Pending
// }

// 16. API Throttling
// If the system has too many requests in a short time period, you may want to queue some transactions to be processed later.
// isSystemOverloaded, err := s.throttlingService.IsSystemOverloaded()
// if err != nil {
// 	s.logger.Error("Failed to check system load", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isSystemOverloaded {
// 	transactionStatus = enum.Queued
// } else {
// 	transactionStatus = enum.Completed
// }

// 17. Users Transaction History
// If a user has a good history of completing transactions, they could be rewarded with faster processing times.
// hasGoodHistory, err := s.userHistoryService.HasGoodTransactionHistory(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to check user's transaction history", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if hasGoodHistory {
// 	transactionStatus = enum.Completed
// } else {
// 	transactionStatus = enum.Pending
// }

// 18. Payment Gateway Availability
// If you rely on a third-party payment gateway, its availability could affect the transaction status.
// isGatewayAvailable, err := s.paymentGatewayService.IsAvailable()
// if err != nil {
// 	s.logger.Error("Failed to check payment gateway availability", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !isGatewayAvailable {
// 	transactionStatus = enum.Pending
// } else {
// 	transactionStatus = enum.Completed
// }

// 19. High-Value Transactions
// Transactions above a certain value might need additional checks or approvals.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if req.Amount > highValueThreshold {
// 	transactionStatus = enum.PendingApproval
// } else {
// 	transactionStatus = enum.Completed
// }

// 20. Currency Exchange
// If the transaction involves currency exchange, the status may depend on successful exchange rate fetching.
// isExchangeRateAvailable, err := s.currencyExchangeService.IsExchangeRateAvailable(req.Currency)
// if err != nil {
// 	s.logger.Error("Failed to check exchange rate availability", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !isExchangeRateAvailable {
// 	transactionStatus = enum.Pending
// } else {
// 	transactionStatus = enum.Completed
// }

// 21. Account Verification Level
// The status of a transaction could be dependent on the verification level of an account.
// verificationLevel, err := s.verificationService.GetAccountVerificationLevel(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to check account verification level", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if verificationLevel < requiredVerificationLevel {
// 	transactionStatus = enum.PendingVerification
// } else {
// 	transactionStatus = enum.Completed
// }

// 22. Geolocation Restrictions
// The transaction may be restricted based on the geographic location of the account holder.
// isLocationAllowed, err := s.geolocationService.IsLocationAllowed(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to check geolocation", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !isLocationAllowed {
// 	transactionStatus = enum.PendingReview
// } else {
// 	transactionStatus = enum.Completed
// }

// 23. Regulatory Compliance
// Transactions might be pending due to compliance with local laws or regulations.
// isCompliant, err := s.complianceService.IsTransactionCompliant(req)
// if err != nil {
// 	s.logger.Error("Failed to check compliance", zap.Error(err))
// 	return res, err
// }

// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !isCompliant {
// 	transactionStatus = enum.PendingReview
// } else {
// 	transactionStatus = enum.Completed
// }

// 24. Seasonal Promotions or Discounts
// The transaction status could be influenced by ongoing promotions or discounts.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if s.promotionService.IsSeasonalPromotionActive() {
// 	transactionStatus = enum.AwaitingPromotionApproval
// } else {
// 	transactionStatus = enum.Completed
// }

// 25. Fraud Detection
// The transaction may be marked as pending if it triggers any fraud detection algorithms.
// isSuspectedFraud, err := s.fraudDetectionService.IsSuspectedFraud(req)
// if err != nil {
// 	s.logger.Error("Failed to run fraud detection", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if isSuspectedFraud {
// 	transactionStatus = enum.PendingReview
// } else {
// 	transactionStatus = enum.Completed
// }

// 26. Business Hours
// Some businesses only process transactions during specific hours.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !s.timeService.IsWithinBusinessHours() {
// 	transactionStatus = enum.PendingBusinessHours
// } else {
// 	transactionStatus = enum.Completed
// }

// 27. Weekend and Holiday Processing
// Transaction status may be dependent on whether it is a weekend or public holiday.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if s.dateService.IsWeekendOrHoliday() {
// 	transactionStatus = enum.PendingWeekendOrHoliday
// } else {
// 	transactionStatus = enum.Completed
// }

// 28. Tier-Based Processing
// Some services have different processing tiers for different levels of service.
// serviceTier, err := s.accountService.GetServiceTier(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to fetch service tier", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if serviceTier == enum.Premium {
// 	transactionStatus = enum.Completed
// } else {
// 	transactionStatus = enum.PendingStandardProcessing
// }

// 29. Age Restrictions
// Certain transactions may have age restrictions.
// isOfLegalAge, err := s.ageVerificationService.IsOfLegalAge(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to verify age", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !isOfLegalAge {
// 	transactionStatus = enum.PendingAgeVerification
// } else {
// 	transactionStatus = enum.Completed
// }

// 30. Operational Flags
// Sometimes, operational flags are used to toggle features or behaviors system-wide.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if s.operationalFlagService.IsFeatureFlagEnabled("NewTransactionModel") {
// 	transactionStatus = enum.PendingNewModelProcessing
// } else {
// 	transactionStatus = enum.Completed
// }


// 31. Geofencing
// Transactions might only be allowed from certain geographical locations.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !s.geoService.IsTransactionLocationAllowed(req.SenderLocation) {
// 	transactionStatus = enum.PendingGeofenceApproval
// } else {
// 	transactionStatus = enum.Completed
// }



// 32. Amount Limits
// Certain transactions might be flagged for manual review if they exceed a particular amount.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if req.Amount > s.transactionService.GetMaxAmountLimit() {
// 	transactionStatus = enum.PendingAmountReview
// } else {
// 	transactionStatus = enum.Completed
// }



// 33. Velocity Checking
// A velocity check can determine if too many transactions are being made in a short period.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if s.transactionService.IsVelocityLimitExceeded(req.SenderAccountID) {
// 	transactionStatus = enum.PendingVelocityReview
// } else {
// 	transactionStatus = enum.Completed
// }



// 34. First-Time Transaction
// First-time transactions might undergo more scrutiny.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if s.transactionService.IsFirstTimeTransaction(req.SenderAccountID) {
// 	transactionStatus = enum.PendingFirstTimeReview
// } else {
// 	transactionStatus = enum.Completed
// }



// 35. Subscription Status
// Some services might only allow transactions for subscribed users.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if !s.subscriptionService.IsSubscribed(req.SenderAccountID) {
// 	transactionStatus = enum.PendingSubscription
// } else {
// 	transactionStatus = enum.Completed
// }


// 36. API Rate Limiting
// Too many API requests from an account in a short period might flag a transaction.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if s.apiService.IsRateLimitExceeded(req.SenderAccountID) {
// 	transactionStatus = enum.PendingRateLimit
// } else {
// 	transactionStatus = enum.Completed
// }


// 37. Account Type Specific Rules
// Different types of accounts (e.g., personal, business) might have different rules.
// accountType, err := s.accountService.GetAccountType(req.SenderAccountID)
// if err != nil {
// 	s.logger.Error("Failed to fetch account type", zap.Error(err))
// 	return res, err
// }
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if accountType == enum.Business {
// 	transactionStatus = enum.PendingBusinessApproval
// } else {
// 	transactionStatus = enum.Completed
// }


// 38. Foreign Transaction
// For international transfers, additional checks may be needed.
// var transactionStatus enum.StatusType // Replace with your actual enum type
// if s.transactionService.IsForeignTransaction(req) {
// 	transactionStatus = enum.PendingForeignTransactionReview
// } else {
// 	transactionStatus = enum.Completed
// }

// *****************************************************************************************************************