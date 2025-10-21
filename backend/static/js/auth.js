// Модуль авторизации
class AuthManager {
    constructor() {
        this.currentUser = null;
        this.init();
    }

    init() {
        // Проверяем, есть ли сохраненные данные пользователя
        this.loadUserFromStorage();
    }

    // Загрузка данных пользователя из localStorage
    loadUserFromStorage() {
        const userData = localStorage.getItem('currentUser');
        if (userData) {
            this.currentUser = JSON.parse(userData);
        }
    }

    // Сохранение данных пользователя в localStorage
    saveUserToStorage(userData) {
        this.currentUser = userData;
        localStorage.setItem('currentUser', JSON.stringify(userData));
    }

    // Проверка авторизации
    isAuthenticated() {
        return this.currentUser !== null;
    }

    // Получение данных текущего пользователя
    getCurrentUser() {
        return this.currentUser;
    }

    // Авторизация пользователя
    login(username, password) {
        // В реальном приложении здесь был бы AJAX запрос к серверу
        // Для демо используем простую проверку
        const validUsers = {
            'admin': { password: 'admin123', name: 'Иванов Александр Сергеевич', role: 'Врач-кардиолог', department: 'Кардиологическое отделение' },
            'doctor': { password: 'doctor123', name: 'Петров Иван Сергеевич', role: 'Врач-терапевт', department: 'Терапевтическое отделение' },
            'nurse': { password: 'nurse123', name: 'Сидорова Мария Петровна', role: 'Медсестра', department: 'Кардиологическое отделение' }
        };

        if (validUsers[username] && validUsers[username].password === password) {
            const userData = {
                username: username,
                name: validUsers[username].name,
                role: validUsers[username].role,
                department: validUsers[username].department,
                loginTime: new Date().toISOString()
            };
            
            this.saveUserToStorage(userData);
            return { success: true, user: userData };
        }
        
        return { success: false, message: 'Неверный логин или пароль' };
    }

    // Выход из системы
    logout() {
        this.currentUser = null;
        localStorage.removeItem('currentUser');
    }

    // Проверка авторизации и перенаправление
    checkAuthAndRedirect() {
        if (!this.isAuthenticated()) {
            // Если пользователь не авторизован, перенаправляем на страницу входа
            window.location.href = '/login';
            return false;
        }
        return true;
    }

    // Обновление данных пользователя на странице
    updateUserInfo() {
        if (!this.currentUser) return;

        // Обновляем имя пользователя в header
        const userNameSpan = document.querySelector('.user-menu span');
        if (userNameSpan) {
            userNameSpan.textContent = this.currentUser.name;
        }

        // Обновляем информацию в welcome banner
        const userRole = document.querySelector('.user-role');
        if (userRole) {
            userRole.textContent = `${this.currentUser.role} | ${this.currentUser.department}`;
        }
    }
}

// Создаем глобальный экземпляр менеджера авторизации
window.authManager = new AuthManager();

// Функция для выхода из системы
function logout() {
    window.authManager.logout();
    window.location.href = '/login';
}

// Функция для проверки авторизации при загрузке страницы
function checkAuthentication() {
    if (!window.authManager.checkAuthAndRedirect()) {
        return;
    }
    
    // Обновляем информацию о пользователе
    window.authManager.updateUserInfo();
}

// Проверяем авторизацию при загрузке DOM
document.addEventListener('DOMContentLoaded', function() {
    // Проверяем, не находимся ли мы на странице входа
    if (window.location.pathname.includes('login')) {
        // Если пользователь уже авторизован, перенаправляем на главную
        if (window.authManager.isAuthenticated()) {
            window.location.href = '/';
        }
        return;
    }
    
    // Для всех остальных страниц проверяем авторизацию
    checkAuthentication();
});
