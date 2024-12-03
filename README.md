# Go Fiber Template v2

A modern, secure, and feature-rich web application template built with Go Fiber, featuring authentication, session management, and a responsive UI.

## 🚀 Features

- **Modern Authentication System**
  - User registration with email verification
  - Secure login with session management
  - Password confirmation and validation
  - "Remember me" functionality
  - Password recovery flow

- **Responsive UI**
  - Built with Bulma CSS framework
  - Mobile-first design approach
  - Modern and clean interface
  - Interactive components with Alpine.js
  - HTMX for dynamic content loading

- **Security Features**
  - Session-based authentication
  - SQLite session storage
  - Secure password hashing
  - HTTP-only cookies
  - CSRF protection
  - Same-site cookie policy

## 🛠️ Tech Stack

- **Backend**
  - Go Fiber v2.52.5
  - GORM with SQLite
  - Go Template Engine
  - Session Management

- **Frontend**
  - Bulma CSS v1.0.2
  - HTMX v2.0.3
  - Alpine.js v3.x
  - Font Awesome v6.4.2

## 📋 Prerequisites

- Go 1.20 or higher
- SQLite3

## 🚀 Quick Start

1. **Use this template**
   - Click the "Use this template" button at the top of this repository
   - Choose "Create a new repository"
   - Fill in your repository name and description
   - Choose public or private visibility
   - Click "Create repository from template"

2. **Clone your new repository**
   ```bash
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Install dependencies**
   ```bash
   go mod download
   ```

5. **Run the application**

   There are two ways to run the application:

   **Option 1: Standard Go Run**
   ```bash
   go run main.go
   ```

   **Option 2: Using Air (Recommended for Development)**
   
   Air provides hot-reloading capabilities, which automatically rebuilds and restarts your application when files change.

   First, install Air:
   ```bash
   # Using go install
   go install github.com/cosmtrek/air@latest

   # Or using curl (Windows PowerShell)
   irm get.scoop.sh | iex
   scoop install air
   ```

   Then run the application with Air:
   ```bash
   # Air will use the .air.toml configuration
   air

   # The application will automatically rebuild when you make changes
   # You'll see output like:
   #   watching .
   #   building...
   #   running...
   ```

   Air configuration (`.air.toml`) includes:
   - Automatic rebuild on file changes
   - Custom build commands
   - Specific file extensions to watch
   - Directories to exclude
   - Custom build output location

## 📁 Project Structure

```
go-fiber-template-v2/
├── app/
│   ├── database/        # Database configurations and models
│   ├── middleware/      # Authentication and session middleware
│   ├── public/
│   │   └── styles/     # CSS and other static assets
│   ├── routes/         # HTTP route handlers
│   └── views/
│       ├── layouts/    # Base layout templates
│       │   └── main.html
│       ├── pages/      # Page templates
│       │   ├── 404.html
│       │   ├── index.html
│       │   ├── login.html
│       │   └── register.html
│       └── partials/   # Reusable template components
│           ├── footer.html
│           ├── head.html
│           └── navbar.html
├── .air.toml           # Air configuration for hot reload
├── .env                # Environment variables
├── .gitignore         # Git ignore rules
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
└── main.go            # Application entry point
```

## 🔒 Security Considerations

- Passwords are hashed using secure algorithms
- Sessions are stored in SQLite database
- Cookies are configured with secure flags
- Input validation on both client and server side
- Protected against CSRF attacks

## 🎨 UI Features

- Responsive navigation with hamburger menu
- Modern authentication forms
- Password strength indicators
- Interactive form validation
- Comprehensive footer with social links
- Loading states and transitions
- Error message handling

## 🔧 Configuration

Key configuration options in `.env`:
```env
PORT: :3000
TITLE: Your App Name
```

## 🚧 Development

1. **Hot Reloading**
   - Uses Air for automatic recompilation
   - Configure in `.air.toml`

2. **Database Migrations**
   - Automatic migrations with GORM
   - Models defined in `app/database`

3. **Template Development**
   - Templates in `app/views`
   - Partials for reusable components
   - Layouts for consistent structure

## 📝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Go Fiber](https://gofiber.io/)
- [Bulma CSS](https://bulma.io/)
- [HTMX](https://htmx.org/)
- [Alpine.js](https://alpinejs.dev/)
- [Font Awesome](https://fontawesome.com/)
