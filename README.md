# CryptoWallet
This project is a modular cryptocurrency trading and management platform implemented in Go.
It leverages several design patterns (such as Command, Factory, and Singleton) to handle user operations, API calls for cryptocurrency pricing, caching, and user management. The project is designed with extensibility in mind, so new data sources or features (like additional listing or filtering options) can be added with minimal changes.
## Features
- **User Management**


  Register and log in users with secure password hashing (using bcrypt).
  A singleton User Database ensures that there is only one global instance managing user data.


- **Cryptocurrency Data Handling**


  Parse JSON responses from cryptocurrency APIs into CryptoCurrency and CryptoCurrencies structures.
  Display listings for individual tokens and collections of tokens.
  Designed to support future expansion (e.g., filtering, sorting, or listing tokens from different sources).


- **Command Pattern Implementation**


   Encapsulate user actions (login, register, buy, sell, add funds, show portfolio, etc.) into separate command objects.
   Factories create commands based on user input, decoupling command creation from execution.
  

- **Price Caching**


  A singleton PriceCache stores cryptocurrency prices, ensuring that data remains fresh for a defined duration and minimizing redundant API calls.

- **API Caller & Price Fetching**


  The ApiCaller package handles HTTP requests to external cryptocurrency APIs.
  Implements a PriceFetcher interface to update prices, allowing for alternative implementations (e.g., file-based or mock updates for testing).


- **Command Engine**


  The Engine package ties everything together with a command loop that processes user input from the console.
  It handles the workflow from login/registration to executing trading commands and displaying portfolio summaries.


- **Utility Functions**


  The helpers package contains utility functions for input validation, output formatting, and database connection management.

## Design Decisions
- **Command Pattern**


  Each user action (login, register, buy, sell, etc.) is encapsulated as a separate command. This decouples command creation and execution, making it easy to add new commands or modify existing ones.

- **Factory Pattern**


  Factories create command objects based on user input. This centralizes the decision-making process for which command to execute, based on application constants and options.

- **Singleton Pattern**


  The User Database (Users) and the Price Cache (PriceCache) are implemented as singletons. This ensures that there is only one global instance of these components throughout the application, simplifying resource management and synchronization.

- **Separation of Concerns**


  The system is divided into multiple packages (users, cryptoCurrency, apiCaller, etc.), each responsible for a specific aspect of the application. This modular approach makes the project easier to maintain and extend.





















