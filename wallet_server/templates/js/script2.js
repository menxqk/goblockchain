$(function () {

    $.ajax({
        url: '/wallet',
        type: 'POST',
        success: function (response) {
            $('#private_key').val(response['private_key']);
            $('#public_key').val(response['public_key']);
            $('#blockchain_address').val(response['blockchain_address']);
            console.info(response);
        },
        error: function (error) {
            console.error(error);
        }
    });

    $('#send_money_button').click(function () {
        let confirm_text = 'Are you sure to send?';
        let confirm_result = confirm(confirm_text);
        if (confirm_result !== true) {
            alert('Canceled');
            return;
        }

        let transaction_data = {
            'sender_private_key': $('#private_key').val(),
            'sender_public_key': $('#public_key').val(),
            'sender_blockchain_address': $('#blockchain_address').val(),
            'recipient_blockchain_address': $('#recipient_blockchain_address').val(),
            'value': $('#send_amount').val(),
        }

        $.ajax({
            url: '/transaction',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(transaction_data),
            success: function (response) {
                console.info(response);
                if (response.message === 'fail') {
                    alert("failed");
                    return;
                }
                alert('Send success');
            },
            error: function (error) {
                console.error(error);
                alert('Send failed');
            }
        });
    });

    function reload_amount() {
        let data = {
            'blockchain_address': $('#blockchain_address').val()
        }
        $.ajax({
            url: '/wallet/amount',
            type: 'GET',
            data: data,
            success: function (response) {
                console.info(response);
                let amount = response['amount'];
                $('#wallet_amount').text(amount);
            },
            error: function (error) {
                console.error(error);
            }
        });
    }

    $('#reload_wallet').click(function () {
        reload_amount();
    });

    setInterval(reload_amount, 10000);

});