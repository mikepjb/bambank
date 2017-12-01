Bambank design -

> I. Top level design

In go

high level requirements: (Sx> specified by doc)
- Sa> log users in to an account
- Sb> view the current balance for the user account
- store account information as a model
- Sc> allow transactions to be made between accounts
- store transaction information as a model
- setup a postgres database to store the data

low level (steps involved to make the high level happen):
log users in to an account -
  - store a string in session
allow transactions to be made between accounts -
  * example: account a transacts 50 euros to account b
  - two transaction objects are created, one for each account
  - account a has a transaction record stored crediting the amount.
  - account b has a transaction record stored debiting the amount.
store <model> information as a model -
  - requires a database connection to be setup

> II. Build strategy

(What steps should I take to build the prototype?)

1. Demonstrate a relationship between the objects in code.
 Relationship 1: accounts and transactions
 * example:
    account_a.transact(account_b, 50)

2. Surface the output of step one via http (Sb)
  - create two accounts: the first is a master account for debitting the free
    money, the other to be a demo account.
  - find an account by a hard-coded id (to be replaced with session string
    lookup)
  - print name for that account
  - print all transaction records under that account

3. Allow changes to the model via calls a.k.a controller methods (Sc)
  ??? detail
  create an account using name, password & password_confirmation
    - auto create transaction crediting the account with 100 euros
  create a transaction between account 123 and 456 for 50 euros

4. User authentication, filter information to their account/transactions only (Sa)
 ??? detail - HOW
 how is this stored? session string most likely, required more explanation

5. Persist model data in postgres
  - use go-pg ORM layer
  - connect
  - create a schema
    - create a table both accounts and transactions
  - create db inserts in model calls via controller methods (instead of creating
    them in memory from step 3)

6. Make it pretty
  ??? detail? have a css party.

7. Host it
  ??? on heroku with free postgres add-on

> III. Model structure

Account
  id
  name
  password
  []transaction_ids // needed or do we query for it against the db?

Transaction
  id
  timestamp
  credit_id
  debit_id
  amount

> IV. References & Go specific questions

how to create a method for a struct?
references: youtube web design golang (personal history)
