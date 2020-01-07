var dz = {
    showFormModal: function (options) {//title,html,button,click
        var defaluts = {
            isBig: false,
            isBiggest: false,
            //loadingbgColorClass: '',
            loadingClassType:1,
            title: '加载中...',
            msg:''
        }, options = $.extend(defaluts, options);
        var modalhtml = "<div class='modal fade " + (options.isBig ? " bs-example-modal-lg" : " bs-example-modal-sm") + "' tabindex='-1' role='dialog' aria-labelledby='mySmallModalLabel'>" +
            "<div class='modal-dialog " + (options.isBig ? (options.isBiggest ? "modal-lg" : " modal-bg") : " modal-sm") + "'>" +
            "<div class='modal-content'>" +
            " <div class='modal-header'>" +
            "<h4 class='modal-title'>" + options.title + "</h4>" +
            "<button type='button' class='close' data-dismiss='modal' aria-label='Close'><span aria-hidden='true'>&times;</span></button>" +
            "</div>" +
            "<div class='modal-body'>" + options.html +
            "</div>" +
            "<div class='modal-footer'>" +
            " <button type='button' class='btn btn-default' data-dismiss='modal'>关闭</button>" +
            "<button type='button' class='btn btn-success'>" + options.button + "</button>" +
            "</div>" +
            "</div></div></div>";

        var modal = $(modalhtml).modal();
        if (options.button == undefined || options.button == "") {
            modal.find(".btn-success").hide();
        }
        modal.find(".btn-success").click(function () {
            if (options.click != undefined) {
                if (options.click() == true) {
                    modal.modal("hide");
                }
            }
        });
        return modal;
    },
    showMsgModal: function (options) {
        var
            html = "<div class='modal fade bs-example-modal-sm' tabindex='-1' role='dialog' aria-labelledby='mySmallModalLabel'>" +
                "<div class='modal-dialog modal-sm'>" +
                "<div class='modal-content'>" +
                " <div class='modal-header'>" +
                "<button type='button' class='close' data-dismiss='modal' aria-label='Close'><span aria-hidden='true'>&times;</span></button>" +
                "<h4 class='modal-title'>" + options.title + "</h4>" +
                "</div>" +
                "<div class='modal-body'>" + options.msg +
                "</div>" +
                "</div></div></div>";

        var modal = $(html).modal()
        return modal
    },
    showErrorMsg: function (msg,container) {
        $("<div class='form-group col-xs-12'><div class='alert alert-danger cw_errormsg' role='alert'><b id='err'>" + msg + "</b></div></div>").appendTo(container);
    },
    showLoading:function(options)
    {
        var modalhtml = "<div class='modal fade bs-example-modal-sm dz-modal-loading' tabindex='-1' role='dialog' aria-labelledby='mySmallModalLabel'>" +
            "<div class='modal-dialog modal-sm'>" +
            "<div class='modal-content'>" +
            " <div class='modal-header'>" +
            "<button type='button' class='close' data-dismiss='modal' aria-label='Close'><span aria-hidden='true'>&times;</span></button>" +
            "<h4 class='modal-title'>" + options.title + "</h4>" +
            "</div>" +
            "<div class='modal-body'>" +
            '<div class="progress">'+
            '<div class="progress-bar progress-bar-striped '+(this.getLoadingColor(options.loadingClassType))+' active" role="progressbar" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100" style="width: 100%">'+
            '<span class="sr-only">100% Complete</span>'+
            '</div>'+
            '</div>'+
            "</div>" +
            "</div></div></div>";
        var modal = $(modalhtml).modal({backdrop:'static'});
        return modal;
    },hideLoading:function(modal){
        $(".dz-modal-loading").modal('hide')
        modal&&	modal.modal("hide")

    }, getLoadingColor: function (type) {
        switch (type) {
            case 1: return "";
            case 2: return "progress-bar-success";
            case 3: return "progress-bar-info";
            case 4: return "progress-bar-warning";
            case 5: return "progress-bar-danger";
        }
        return "";
    }
}