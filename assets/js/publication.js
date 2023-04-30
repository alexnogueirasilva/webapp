$('#form-publication').on('submit', createPublication);

function createPublication(){
    event.preventDefault();

    let title = $('#title').val();
    let content = $('#content').val();

    $.ajax({
        url: '/publications',
        method: 'POST',
        data: {
            title: title,
            content: content
        }
    }).done(function() {
        window.location.href = '/home';
    }).fail(function(error) {
        alert('Erro ao criar publicacao')
        console.log(error);
    });
}