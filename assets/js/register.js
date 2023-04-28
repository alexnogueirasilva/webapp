$('#register-user').on('submit', registerUser);

function registerUser(event) {
    event.preventDefault();

    let name = $('#name').val();
    let email = $('#email').val();
    let nickname = $('#nickname').val();
    let password = $('#password').val();
    let confirmPassword = $('#confirm-password').val();

    if(password !== confirmPassword){
        alert('Password do not match');
        return;
    }

    $.ajax({
        url: '/users',
        method: 'POST',
        data: {
            name: name,
            email: email,
            nickname: nickname,
            password: password
        }
    }).done(function (response) {
        alert('User registered successfully');
        // window.location.href = '/login';
    }).fail(function (error) {
        alert('Something went wrong');
        console.log(error);
    });

}