
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