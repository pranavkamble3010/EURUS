$(document).ready(function(){

    $.get("http://localhost:8080/v1/interactions/discussion", function(data, status){
      if(status == "success")
      {
          discussions = data; 
          renderDiscussionsData(discussions);
      }
        
      if(status == "error")
        alert("Error: " + status);
    });
}

);   //$(document).ready ends


function renderDiscussionsData(discussions)
{
    for(var i=0;i<discussions.length;i++)
    {
        console.log(discussions[i]);
        // ques = $("<div id="+ discussions[i].ObjectId +"></div>").text(discussions[i].Description).addClass("bg-light m-2 p-2 py-2 bd-highlight border border-primary");
    
        // ques.appendTo("#discussionsDisplay");
        card = $("<div id=\"card"+i+"\"></div>").addClass("card");
        cardHeader = $("<div id=cardHeader"+i+"></div>").addClass("card-header");
        h2 = $("<h2 id=h2disc"+i+"></h2>").addClass("mb-0");
        discussionText = $("<div id=\"discussionText_"+discussions[i].ObjectId+"\"><div>").text(discussions[i].Description).addClass("text-justify pb-3");
        button = $("<button type=\"button\" id=\""+ discussions[i].ObjectId +"\" data-toggle=\"collapse\" data-target=\"#collapse"+discussions[i].ObjectId+"\" aria-expanded=\"true\" aria-controls=\"collapseOne\"></button>")
        .addClass("btn btn-link text-left font-italic").text("Click here to view all responses..");

        /************Response section start*****************/
        formResponse = $("<div class=\"input-group input-group-sm mb-3\"></div>");
        txtResponse = $("<input type=\"text\" id=\"user_"+localStorage.getItem("userObjectID")+"_respTo_"+discussions[i].ObjectId+"\" placeholder=\"Write your response\">")
                        .addClass("form-control");

        
        btnPostResponse = $("<button type=\"submit\" id=\""+discussions[i].ObjectId+"\">POST</button>").addClass("btn btn-outline-secondary");
        //btnPostResponse.on('click',function(){alert(this.id)});
        btnPostResponse.on('click',respondDisc);
        btnPostResponseWrapper = $("<div></div>").addClass("input-group-append");
        btnPostResponse.appendTo(btnPostResponseWrapper);
        
        txtResponse.appendTo(formResponse);
        btnPostResponseWrapper.appendTo(formResponse);
        /************Response section end*****************/

        //get all the responses for this interaction ID when clicked
        button.on("click",function(){
            if($("#cardBody_"+this.id).text() == ""){
                $.get("http://localhost:8080/v1/responses/intrresp/"+this.id, function(data, status){
                responses = data;
                if(responses != null)
                {
                for(var j=0;j<responses.length;j++)
                {
                    answer = $("<div id=\""+responses[j].InteractionId+"_response_"+j+"\"></div>")
                    .addClass("p-2 bd-highlight border border-primary rounded");
                    answer.text(responses[j].ResponseContent);
                    answer.appendTo("#cardBody_"+responses[j].InteractionId);
                    //console.log(responses[j]);
                }
                }     
            });
            }
        })
        cardBody = $("<div id=\"collapse"+discussions[i].ObjectId+"\" aria-labelledby=\"cardHeader"+i+"\" data-parent=\"#discussionsDisplay\"></div>").addClass("collapse");
        cardBodyContent=$("<div id=\"cardBody_"+discussions[i].ObjectId+"\"></div>").addClass("card-body");
        cardBodyContent.appendTo(cardBody);

        discussionText.appendTo(cardHeader);
        button.appendTo(h2);
        
        formResponse.appendTo(cardHeader);
        h2.appendTo(cardHeader);

        cardHeader.appendTo(card);
        cardBody.appendTo(card);
        
        card.appendTo("#discussionsDisplay");
        
    }
}

function respondDisc(){
    cardId = this.id;
    //responseCardId = "#cardBody_"+cardId;
    responseText = $("#user_"+localStorage.getItem("userObjectID")+"_respTo_"+this.id).val();

    inputJson = '{"InteractionId": "'+ this.id +
                '", "ResponseType":0,'+ 
                '"OwnerId":"'+localStorage.getItem("userObjectID") +'",'+
                '"ResponseContent":"'+responseText+'"}';

    //alert(inputJson);
    $.post("http://localhost:8080/v1/responses",inputJson,function(data,status){

    if(status == 'success')
    {
        //alert(data);
        $("#cardBody_"+cardId).html("");
        //populateResponses();
        $.get("http://localhost:8080/v1/responses/intrresp/"+cardId, function(data, status){
            responses = data;
            if(responses != null)
            {
            for(var j=0;j<responses.length;j++)
            {
                answer = $("<div id=\""+responses[j].InteractionId+"_response_"+j+"\"></div>")
                .addClass("p-2 bd-highlight border border-primary rounded");
                answer.text(responses[j].ResponseContent);
                answer.appendTo("#cardBody_"+responses[j].InteractionId);
                //console.log(responses[j]);
            }
            }     
        })
    }
    else{
        alert('failed!');
    }

    });

}