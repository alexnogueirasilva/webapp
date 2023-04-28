$('#login').on('submit',  login);

function login(event) {
    event.preventDefault();

    let email = $('#email').val();
    let password = $('#password').val();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email: email,
            password: password
        }
    }).done(function() {
        window.location.href = '/home';
    }).fail(function(error) {
        console.log(error);
    });
}