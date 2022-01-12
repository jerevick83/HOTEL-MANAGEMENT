let attention = prompt();
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
})()

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
        const {value: formValues} = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
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
                const calendarPicker = document.querySelector(".reservation-dates-modal");
                new DateRangePicker(calendarPicker, {
                    format: "yyyy-mm-dd",
                    orientation: "top",
                    todayBtn: true,
                    todayBtnMode: 1,
                    autohide: true,
                    allowOneSidedRange: true,
                    showOnFocus: true
                });
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            },
            didOpen: () => {
                document.getElementById('start').removeAttribute("disabled");
                document.getElementById('end').removeAttribute("disabled");
            }
        })

        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
    }
    return ({
            success,
            warning,
            error,
            searchAvail
        }
    )
}