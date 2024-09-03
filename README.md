# GO_Projects

## Ideas

###  1. TodoApp
        Commandline Todo app, following requirements
        - Commands: add, viewList, viewTask, completeTask
        - data store = csv
        - potential pkgs: tabwriter, timediff, cobra
###  2. API
        StatelessAPI; Calculator service
        -ADD,Subtract,Divied,times
###  3. Create a student management system
        A student management system allows school administrators to manage student enrollments and generate reports. Administrators should be able to perform CRUD operations to work with student enrollment, set grades, schedule classes, and generate reports.
        This project is an interactive menu-driven command-line program. The options available for a user should use a numbering system, and users can type the number of the option they wish to access.
        For example
        Teachers should have an "Enter Grades" menu option, which allows them to enter grades for students enrolled in their classes.
        This project has a lot of small parts, so splitting it into multiple layers will help manage the project's complexity:
        Data layer: handles database queries
        Business layer (optional): verifies that data going in and out of the database is correct. An example would be verifying if a user has permission to view certain things, or validating information before saving it to the database.
        View layer: displays information to the user
        Interaction layer: connects the view layer with the business or data layer. This layer contains all possible actions that your application can perform.
        To get started, you'll need to:
        Design the database tables
        Write database queries using raw SQL or with an ORM such as GORM
        Design a menu system
        Write functions to select a menu option
        Create input parsing functions
        Output formatted reports
        Level up even further!
        Level up your skills by implementing more features:
        Log all changes to an audit log in the database
        Create separate permissions for administrators, instructors, and students
        Allow or deny access to specified portions of the system based on the user's position (listed above)
        Add a web interface which is separate from the command line interface. Both should use the same database
###  4. File Encryption Tool
        Develop a program that encrypts and decrypts files using symmetric encryption algorithms like AES. This project teaches beginners about file handling, encryption techniques, and error handling in Go.
        What You’ll Learn From This Project Idea:
        File handling operations in Go
        Understanding encryption algorithms (e.g., AES)
        Error handling for file operations and encryption
        Command-line interface for user interaction
### 5. Blockchain Implementation
        Build a blockchain from scratch, including features like proof-of-work consensus, transaction validation, and peer-to-peer networking. This project delves into cryptography, distributed systems, and decentralized application development using Go.
        What You’ll Learn From This Project Idea:
        Cryptography fundamentals and blockchain concepts
        Peer-to-peer networking and decentralized consensus mechanisms
        Smart contract development and decentralized application (dApp) deployment





$ curl -d '{"num1":2, "num2":3}' http://localhost:8080/add/
