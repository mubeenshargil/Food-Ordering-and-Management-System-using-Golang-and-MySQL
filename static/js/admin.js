$(document).ready(function(){

    $('#menuTable').DataTable({

        pageLength: 5,

        lengthMenu: [
            [5,10,25,50],
            [5,10,25,50]
        ],

        ordering: true,

        searching: true,

        responsive: true

    });

});

$(document).ready(function(){

    var table = $('#ordersTable').DataTable({

        pageLength: 5,

        lengthMenu: [
            [5,10,25,50],
            [5,10,25,50]
        ]

    });

    $('#statusFilter').on('change', function(){

        table.column(6)
             .search(this.value)
             .draw();

    });

});


$(document).ready(function(){

    $('#usersTable2').DataTable({

        pageLength: 5,

        lengthMenu: [
            [5,10,25,50],
            [5,10,25,50]
        ]

    });

});

$('.reportTab').DataTable({

        pageLength: 5,

        lengthMenu: [
            [5,10,25,50],
            [5,10,25,50]
        ],

        ordering: true,

        searching: true,

        responsive: true

    });

const status = document.getElementById("status");
const reasonBox = document.getElementById("cancelReasonBox");
const reason = document.getElementById("cancel_reason");
const btn = document.getElementById("updateBtn");

status.addEventListener("change", function() {

    if (this.value === "CANCELLED") {

        reasonBox.style.display = "block";
        reason.required = true;

        btn.disabled = true;

    } else {

        reasonBox.style.display = "none";
        reason.required = false;
        reason.value = "";

        btn.disabled = false;
    }
});

reason.addEventListener("input", function() {

    if (status.value === "CANCELLED") {

        btn.disabled = this.value.trim() === "";
    }
});