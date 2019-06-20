$(document).ready(function(){

    //$('#ask').hide();
    
    localStorage.setItem("userObjectID","5ca6ad2c58fb5828cc0b3cb2");    //temporary setting

    $("#login").on('click',function(){

        email = $('#email_id').val();
        password = $('#password').val();

        inputJson = '{"Email" :"'+ email +'", "Password":"'+ password +'"}'

        $.post("http://localhost:8080/v1/user",inputJson,login);

        //alert(username+" "+password);

        function login(data,status)
        {
            if(status == "success")
            {
                $('#ask').show();
                localStorage.setItem("email",email);
                console.log(data);
            }
            else{
                alert("Error occured!");
            }
           

            //alert(localStorage.getItem("email"));
        }
    });


    $("#btnSearch").on('click',function(){

        var searchString = $("#txtSearch").val();

        inputJson = '{"Tags" :"'+ searchString +'\"}'

        $.post("http://localhost:8080/v1/interactions/search/",inputJson,extractResult);


        function extractResult(data,status)
        {
            if(status == "success")
            {
                //console.log(data);
                $("#discussionsDisplay").html("")   //empty current data
                $("#questionsDisplay").html("")    //empty current data
                $("#debatesDisplay").html("")    //empty current data
                searchResult = data;
                debates = [];
                questions = [];
                discussions = [];
                for(var i=0;i<searchResult.length;i++)
                {
                    if(searchResult[i].InteractionType == 1)
                    {
                        
                        debates.push(searchResult[i]);
                    }

                    else if(searchResult[i].InteractionType == 0)
                    {
                        
                        questions.push(searchResult[i]);
                    }

                    else if(searchResult[i].InteractionType == 2)
                    {
                        
                        discussions.push(searchResult[i]);
                    }
                }

                renderDebateData(debates);
                renderQuestionsData(questions);
                renderDiscussionsData(discussions);
            }
        }

    });

    $("#btnPost").on('click',function(){

        topic = $("#txtopic").val();
        desc = $("#txtDescription").val();
        type = $("#selectType").val();
        tags = $("#txtags").val();

        inputJson = '{"InteractionType":'+ type +','+
                        '"Topic" :"'+ topic +'",'+
                        '"Tags":"'+ tags +'",'+
                        '"OwnerId":"5ca6ad2c58fb5828cc0b3cb2",'+
                        '"Description":"'+ desc +'"'+
                    '}'

        $.post("http://localhost:8080/v1/interactions/",inputJson,function(data,status){
            
        console.log(data);
    
    });
    })


})