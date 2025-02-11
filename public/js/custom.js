const csrf_token = $('meta[name="csrf-token"]').attr('content');
function disabledButton(id,status) {
    const $button = $(`button[id=${id}]`)
    if (status) {
        $button.data('original-text', $button.text());
        $button.text("Loading...");
        $button.prop('disabled', true);
    } else {
        $button.prop('disabled', false);
        $button.text($button.data('original-text'));
    }
}

// URL: example: /ajax/auth/login,
//     IdForm: example: loginForm,
//     IdButton: Button submit example: login.
// Success: function(dataDoneParse)
// Error: function(jqXHR,dataDoneParse)
// before: function(req)
function request(url, idForm, idButton, success = function(){}, err = function(){},before = function(){}) {

    $.ajax({
        url: url,
        method: "POST",
        data: $(`form[id=${idForm}]`).serialize(),
        headers: {
            "X-Csrf-Token": csrf_token,
        },
        beforeSend: function (req) {
            disabledButton(idButton,true)
            before(req)
        },
        success: function (data) {
            disabledButton(idButton,false)
            let json;
            try {
                json = JSON.parse(data);
            } catch (e) {
                console.log("Error: " + e);
                json = data;
            }

            Swal.fire({
                title: json.Title,
                icon: json.Icon,
                text: json.Body
            })

            success(json);
        },
        error: function (jqXHR) {
            disabledButton(idButton, false);
            if ( jqXHR.status === 403 || jqXHR.status === "403") {
                Swal.fire({
                    title: "Gagal",
                    icon: "info",
                    text: "Anda tidak memiliki akses"
                })
            } else {
                const data = JSON.parse(jqXHR.responseText);
                Swal.fire({
                    title: data.Title,
                    icon: data.Icon,
                    text: data.Body
                })
                err(jqXHR,data);
            }
        }

    })
}

