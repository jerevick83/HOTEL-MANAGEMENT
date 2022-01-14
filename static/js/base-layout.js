(function () {
    'use strict';

    window.addEventListener('load', () => {
        const forms = document.querySelectorAll('.needs-validation');
        // Loop over the forms and prevent submission
        Array.prototype.filter.call(forms, (form) => {
            form.addEventListener('submit', function (event) {
                if (form.checkValidity === false) {
                    event.preventDefault()
                    event.stopImmediatePropagation()
                }
                form.classList.add('was-validated')
            }, false)
        })
    })
})();

let attention = prompt();
function prompt() {
    let success = (c) => {

        const {
            msg = "",
            icon = "success",
            position = "top-end"
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            icon: icon,
            position: position,
            showCancelButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    };
    let warning = () => {
        console.log("Hello From warning")
    };
    let error = () => {
        console.log("This is an error")
    };

    // searchAvail check for room availability
    let searchAvail = async (c) => {
        const {
            msg = "",
            title = ""
        } = c
        const {value: result} = await Swal.fire({
            title: title,
            html: msg,
            backdrop: true,
            showCancelButton: true,
            customClass: {
                cancelButton: "cancelBtn",
                confirmButton: "confirmButton"
            },
            focusConfirm: false,
            width: '35rem',
            scrollbarPadding: false,
            inputAutoTrim: true,
            buttonsStyling: false,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen()
                }
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            },

        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result.value !== "") {
                    if (c.callback !== undefined) {
                        c.callback(result)
                    }
                } else {
                    c.callback(false)
                }
            } else {
                c.callback(false)
            }
        }
    }
    return {
        success,
        warning,
        error,
        searchAvail
    }
}


